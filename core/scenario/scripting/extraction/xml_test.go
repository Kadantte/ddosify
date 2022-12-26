package extraction

import (
	"fmt"
	"strings"
	"testing"
)

func TestXmlExtraction(t *testing.T) {
	expected := "XML Title"
	xmlSource := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" ?>
	<rss version="2.0">
	<channel>
	  <item>
		<title>%s</title>
	  </item>
	</channel>
	</rss>`, expected)

	xe := XmlExtractor{}
	xpath := "//item/title"
	val, err := xe.extractFromByteSlice([]byte(xmlSource), xpath)

	if err != nil {
		t.Errorf("TestXmlExtraction %v", err)
	}

	if !strings.EqualFold(val.(string), expected) {
		t.Errorf("TestXmlExtraction expected: %s, got: %s", expected, val)
	}
}

func TestXmlExtraction_PathNotFound(t *testing.T) {
	expected := "XML Title"
	xmlSource := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" ?>
	<rss version="2.0">
	<channel>
	  <item>
		<title>%s</title>
	  </item>
	</channel>
	</rss>`, expected)

	xe := XmlExtractor{}
	xpath := "//item3/title"
	_, err := xe.extractFromByteSlice([]byte(xmlSource), xpath)

	if err == nil {
		t.Errorf("TestXmlExtraction_PathNotFound, should be err, got :%v", err)
	}
}
