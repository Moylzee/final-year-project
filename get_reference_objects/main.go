package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var (
	AllObjects         []string
	NewSwaggerFilePath = "../bucket/latest_swagger/latest_swagger.json"
)

func main() {
	// Read Swagger File
	swagger := readSwaggerFile()

	// Get All Reference Process
	AllRefs(swagger)

	log.Println("Writing All Objects to File")
	if err := WriteToFile(AllObjects, "all_refs.json"); err != nil {
		log.Printf("Error writing to file: %v", err)
	}
	log.Println("All Objects Written to File")

	log.Println("Found All Objects used by CX as Code - Adding Them to JSON")

	definitionMap := make(map[string]interface{})
	for _, object := range AllObjects {
		definition, exists := swagger[object]
		if !exists {
			continue
		}
		definitionMap[object] = definition
	}

	log.Println("Writing CX as Code Objects to Swagger File")
	if err := WriteToFile(definitionMap, NewSwaggerFilePath); err != nil {
		log.Printf("Error writing to swagger file: %v", err)
	}
	log.Println("CX as Code Objects Written to Swagger File")
}

func WriteToFile(d interface{}, filePath string) error {
	data, err := json.MarshalIndent(d, "", "  ")

	if err = os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}
	return nil
}

func readSwaggerFile() map[string]interface{} {
	log.Println("Reading Swagger File From File")

	data, err := os.ReadFile(NewSwaggerFilePath)
	if err != nil {
		log.Printf("Error reading swagger file: %v", err)
	}

	var swagger map[string]interface{}
	if err = json.Unmarshal(data, &swagger); err != nil {
		log.Printf("Error parsing swagger JSON: %v", err)
	}

	log.Println("Read Swagger File From File")
	return swagger
}

func AllRefs(swagger map[string]interface{}) {
	for _, ref := range CxObjects {
		log.Printf("Finding Reference Objects for %s", ref)
		FindAllRefs(swagger[ref])
		if !Contains(AllObjects, ref) {
			AllObjects = append(AllObjects, ref)
		}
	}
}

func FindAllRefs(schema interface{}) {
	if schemaMap, ok := schema.(map[string]interface{}); ok {
		for key, value := range schemaMap {
			if strings.HasSuffix(key, "$ref") {
				value = strings.TrimPrefix(value.(string), "#/definitions/")

				if !Contains(AllObjects, value.(string)) {
					AllObjects = append(AllObjects, value.(string))
				}
			}
			FindAllRefs(value)
		}
	}
}

func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

var CxObjects = []string{
	"AuthzDivision", "DataTable", "EmergencyGroup",
	"Grammar", "GrammarLanguage",
	"IVR", "ScheduleGroup", "Schedule", "Prompt",
	"DomainOrganizationRoleCreate",
	"InstagramIntegrationRequest", "MessagingSettingRequest",
	"SupportedContent", "ExternalMetricDefinitionCreateRequest",
	"ExternalContact", "FlowLogLevelRequest", "FlowMilestone",
	"FlowOutcome", "GroupCreate", "RoleDivisionGrants",
	"CreateIntegrationRequest", "PostActionInput", "Credential",
	"PublishDraftInput", "FacebookIntegrationRequest",
	"OutcomePredictorRequest", "JourneyView", "KnowledgeDocumentReq",
	"LocationCreateDefinition", "TrustRequestCreate",
	"AttemptLimits", "CallableTimeSet", "ResponseSet",
	"Campaign", "CampaignRule", "ContactList", "ContactListTemplate",
	"ContactListFilter", "DigitalRuleSet", "DncListCreate",
	"FileSpecificationTemplate", "RuleSet", "CampaignSequence",
	"CreateTriggerRequest", "PolicyCreate", "Library", "Response",
	"CreateResponseAssetRequest", "InboundDomain", "InboundRoute",
	"Language", "CreateQueueRequest", "RoutingSkill",
	"SkillGroupWithMemberDivisions", "SmsAddressProvision",
	"CreateUtilizationLabelRequest", "WrapupCodeRequest",
	"PublishScriptRequestData", "WorkbinCreate", "WorkitemCreate",
	"WorkitemStatusCreate", "Team", "DIDPool", "EdgeGroup",
	"ExtensionPool", "Phone", "PhoneBase", "OutboundRouteBase",
	"Site", "CreateUser",
	"WebDeploymentConfigurationVersion", "WebDeployment",
}
