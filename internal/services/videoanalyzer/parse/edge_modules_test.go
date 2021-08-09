package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/resourceid"
)

var _ resourceid.Formatter = EdgeModulesId{}

func TestEdgeModulesIDFormatter(t *testing.T) {
	actual := NewEdgeModulesID("12345678-1234-9876-4563-123456789012", "resGroup1", "analyzer1", "edgemodule1").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/videoAnalyzers/analyzer1/edgeModules/edgemodule1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestEdgeModulesID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *EdgeModulesId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/",
			Error: true,
		},

		{
			// missing VideoAnalyzerName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/",
			Error: true,
		},

		{
			// missing value for VideoAnalyzerName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/videoAnalyzers/",
			Error: true,
		},

		{
			// missing EdgeModuleName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/videoAnalyzers/analyzer1/",
			Error: true,
		},

		{
			// missing value for EdgeModuleName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/videoAnalyzers/analyzer1/edgeModules/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Media/videoAnalyzers/analyzer1/edgeModules/edgemodule1",
			Expected: &EdgeModulesId{
				SubscriptionId:    "12345678-1234-9876-4563-123456789012",
				ResourceGroup:     "resGroup1",
				VideoAnalyzerName: "analyzer1",
				EdgeModuleName:    "edgemodule1",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/RESOURCEGROUPS/RESGROUP1/PROVIDERS/MICROSOFT.MEDIA/VIDEOANALYZERS/ANALYZER1/EDGEMODULES/EDGEMODULE1",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := EdgeModulesID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.VideoAnalyzerName != v.Expected.VideoAnalyzerName {
			t.Fatalf("Expected %q but got %q for VideoAnalyzerName", v.Expected.VideoAnalyzerName, actual.VideoAnalyzerName)
		}
		if actual.EdgeModuleName != v.Expected.EdgeModuleName {
			t.Fatalf("Expected %q but got %q for EdgeModuleName", v.Expected.EdgeModuleName, actual.EdgeModuleName)
		}
	}
}
