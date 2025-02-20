import pandas as pd

resultsFilepath = "../../bucket/comparison_results/results.csv"

final = pd.DataFrame(columns=['Attribute', 'Status'])

data = pd.read_csv(resultsFilepath)

print(data)
print("------")
print(data['Attribute'])

object_name = data.get('Attribute')[0].split('.')[0]
print(object_name)

# Check if every row starting with object_name has status added and Anchor NaN
matching_rows = data[data['Attribute'].str.startswith(object_name)]
has_status_added_and_anchor_nan = (
    (matching_rows['Status'] == 'added') &
    (matching_rows['Anchor Value'].isna().all())
)

if has_status_added_and_anchor_nan.all():
    print(f"All {len(matching_rows)} rows starting with '{object_name}' have Status 'added' and Anchor NaN")
    final['Attribute'] = object_name
    final['Status'] = 'new object'
else:
    print(f"Not all rows starting with '{object_name}' have Status 'added' and Anchor NaN'")
    
    # If we want to see which ones don't meet the condition
    non_meeting_rows = matching_rows[~(matching_rows['Status'] == 'added') | (matching_rows['Anchor'].notna())]
    print("\nRows not meeting the condition:")
    print(non_meeting_rows)
