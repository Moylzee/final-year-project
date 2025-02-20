package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	SwaggerUrl            = "https://s3.dualstack.us-east-1.amazonaws.com/inin-prod-api/us-east-1/public-api-v2/swagger-schema/publicapi-v2-latest.json"
	NewSwaggerFilePath    = "../Bucket/latest_swagger/latest_swagger.json"
	AnchorSwaggerFilePath = "../Bucket/anchor_swagger/anchor.json"
	AllObjects            []string
)

func main() {
	var data map[string]interface{}

	log.Println("Retrieving Swagger File From URL")
	resp, err := http.Get(SwaggerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Failed to retrieve Swagger File")
	}

	log.Println("Retrieved Swagger File")	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	// Filter By CX as Code Definitions
	defJson := data["definitions"].(map[string]interface{})

	// Write CX Definitions to File
	log.Println("Writing Swagger To File")
	WriteToFile(defJson, NewSwaggerFilePath)
	log.Printf("Swagger Written To File: %s", NewSwaggerFilePath)
}

func WriteToFile(d interface{}, filePath string) {
	data, err := json.MarshalIndent(d, "", "  ")

	if err = os.WriteFile(filePath, data, 0644); err != nil {
		log.Fatal(err)
	}
}

var CxObjects = []string{
	"DataTable", "EmergencyGroup", "Grammar", "GrammarLanguage",
	"IVR", "ScheduleGroup", "Schedule", "Prompt",
	"AuthzDivision", "DomainOrganizationRoleCreate",
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
	"Site", "CreateUser", "WebDeploymentConfigurationVersion",
	"WebDeployment",
}
