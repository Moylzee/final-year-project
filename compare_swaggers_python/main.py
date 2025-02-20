import json
from matplotlib import pyplot as plt
import pandas as pd

LATEST_SWAGGER_FILEPATH = "../Bucket/latest_swagger/latest_swagger.json"
ANCHOR_SWAGGER_FILEPATH = "../bucket/anchor_swagger/anchor.json"
DESTINATION_FILEPATH = "../Bucket/comparison_results/results.json"

# Function to get enum values from a JSON object
def get_enum_values(json_obj, key_prefix):
    enum_values = []
    for key, value in json_obj.items():
        if key.startswith(key_prefix) and ".enum" in key:
            enum_values.append(value)
    return enum_values

def compare_enums(enum_values_anchor, enum_values_swagger):
    added = list(set(enum_values_swagger) - set(enum_values_anchor))
    removed = list(set(enum_values_anchor) - set(enum_values_swagger))
    return added, removed

def remove_properties(diff):
    new_diff = {}  # Create a new dictionary to store the modified structure
    
    for key, value in diff.items():
        if isinstance(value, dict):
            # Handle nested dictionaries
            new_value = remove_properties(value)
            
            # Modify the key if necessary
            new_key = key.replace(".properties.", ".") if ".properties." in key else key
            new_diff[new_key] = new_value
        else:
            # Modify the key if necessary
            new_key = key.replace(".properties.", ".") if ".properties." in key else key
            new_diff[new_key] = value
            
    return new_diff

# Function to compare two flattened JSON objects with naming convention handling
def compare_jsons(anchor, swagger):
    diff = {}
    all_keys = set(anchor.keys()).union(set(swagger.keys()))

    for key in all_keys:
        value1 = anchor.get(key)
        value2 = swagger.get(key)

        # Skip Attributes that have selfUri 
        if ".selfUri." in key:
            continue

        # Handle Enums Appropriately 
        if ".enum" in key:
            enum_values_anchor = get_enum_values(anchor, key.split(".enum")[0])
            enum_values_swagger = get_enum_values(swagger, key.split(".enum")[0])
            added, removed = compare_enums(enum_values_anchor, enum_values_swagger)
            if added and removed:
                diff[key.split(".enum")[0]] = {
                    'added': added,
                    'removed': removed,
                    'new_enum': enum_values_swagger,
                }
            elif added and not removed:
                diff[key.split(".enum")[0]] = {
                    'added': added,
                    'new_enum': enum_values_swagger,
                }

            elif not added and removed:
                diff[key.split(".enum")[0]] = {
                    'removed': removed,
                    'new_enum': enum_values_swagger,
                }
            continue

        if ".description" in key and ".type" not in key:
            if value1 != value2:
                diff[key] = {
                    'old_description': value1,
                    'new_description': value2,
                }
            continue

        if key.endswith(".readOnly"):
            if value1 != value2:
                diff[key] = {
                    'old_readOnly': value1,
                    'new_readOnly': value2
                }
            continue        

        if key.endswith(".example"):
            if value1 != value2:
                diff[key] = {
                    'old_example': value1,
                    'new_example': value2
                }
            continue

        if key.endswith(".type"):
            if ".type" in key and len(key.split(".")) == 2:
                if key not in anchor:
                    diff[key] = {
                        'new_type': value2
                    }
            elif value1 != value2:
                diff[key] = {
                    'old_type': value1,
                    'new_type': value2
                }
            continue

    diff = remove_properties(diff)
    return diff

def new_object_diff(diff):
    new_diff = {}
    
    for key, value in list(diff.items()):  # Iterate over a copy of diff items to allow for modification
        if key.endswith(".type") and len(key.split(".")) == 2:
            base_object = key.split(".")[0]
            
            # Ensure new object exists in new_diff before adding attributes
            if base_object not in new_diff:
                new_diff[base_object] = {}

            new_diff = add_attributes_to_new_object(base_object, diff, new_diff)

            # After adding the attributes to new_diff, remove them from diff
            if base_object in diff:
                del diff[base_object]

    print(f"New Diff: {new_diff}")
    return new_diff

def add_attributes_to_new_object(base_object, diff, new_diff):
    for key, value in diff.items():
        if key.startswith(base_object + ".") and key != base_object + ".type":
            key = key.replace(base_object + ".", "")
            new_diff[base_object][key] = {}

            # Copy only relevant attributes (exclude `old_` values)
            for k, v in value.items():
                if not k.startswith("old_") or v is not None:
                    new_diff[base_object][key][k] = v
            
            # Remove key if it's empty after filtering
            if not new_diff[base_object][key]:
                del new_diff[base_object][key]
    return new_diff

# TODO: It Adds the New Objects fine but does not add the objects that have been modified 

# Load JSON files
print("Loading JSON files...")
anchor = json.load(open(ANCHOR_SWAGGER_FILEPATH))
latest_swagger = json.load(open(LATEST_SWAGGER_FILEPATH))
print("JSON files loaded")

# Compare the flattened JSON objects
print("Comparing JSON files...")
differences = compare_jsons(anchor, latest_swagger)
differences = new_object_diff(differences)
print("Comparison complete")

if len(differences) == 0:
    print("No differences found.")
    with open(DESTINATION_FILEPATH, 'w') as f:
        f.write("")
    exit(0)

# Save results to JSON
print("Saving results to JSON...")
with open(DESTINATION_FILEPATH, 'w') as f:
    json.dump(differences, f, indent=4, sort_keys=True)
print("Results saved to JSON")