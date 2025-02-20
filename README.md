## bucket 
Used to Store any files which will be accessed by multiple Scripts
### Latest Swagger 
Stores the latest CX as Code Definitions from the swagger
### Anchor Swagger 
Stores the Swagger Definitions Previously used as the latest
The Definitions to use to check if there have been updates

## get_new_swagger
This Script is used to handle the retrieval of the latest swagger from the url

1) Run the get_new_swagger to grab the latest swagger
2) Run the get_reference_objects so we only check relevant schemas
3) Flatten the JSON for easier Comparison ??
4) Compare the two jsons
5) Output the differences to a CSV file
6) Overwrite the anchor.json file (With just cx definitions?)