package fakedata

import (
	"cloudsign/fakeserver/models"
	"testing"
)

func TestGenerateHyphenTruncatedUUIDTagName(t *testing.T) {
	type Model struct {
		ID string `json:"id,omitempty" faker:"hyphenTruncatedUUID"`
	}

	fakemodel, err := GenerateFakeData(Model{})
	if err != nil {
		t.Fatal(err)
	}

	if fakemodel.(Model).ID == "" {
		t.Fatal("id is empty")
	}

	t.Log(fakemodel.(Model).ID)
}

func TestRFC3339TagName(t *testing.T) {
	type Model struct {
		DateString string `json:"dateString,omitempty" faker:"rfc3339"`
	}

	fakemodel, err := GenerateFakeData(Model{})
	if err != nil {
		t.Fatal(err)
	}

	if fakemodel.(Model).DateString == "" {
		t.Fatal("id is empty")
	}

	t.Log(fakemodel.(Model).DateString)
}

func TestDocumentModel(t *testing.T) {
	fakemodel, err := GenerateFakeData(models.DocumentModel{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(fakemodel)
}

func TestDocumentListModel(t *testing.T) {
	fakemodel, err := GenerateFakeData(models.DocumentListModel{})
	if err != nil {
		t.Fatal(err)
	}

	if fakemodel.(models.DocumentListModel).Total != 5 {
		t.Fatalf("total is not equal to 5, actual %v", fakemodel.(models.DocumentListModel).Total)
	}

	if fakemodel.(models.DocumentListModel).Page != 1 {
		t.Fatalf("page is not equal to 1, actual %v", fakemodel.(models.DocumentListModel).Page)
	}

	documentCount := len(fakemodel.(models.DocumentListModel).Documents)
	if documentCount < 1 || documentCount > 3 {
		t.Fatalf("document length is not allowed range, actual %v", documentCount)
	}
}
