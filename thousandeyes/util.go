package thousandeyes

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/william20111/go-thousandeyes"
)

func expandAgents(v interface{}) thousandeyes.Agents {
	var agents thousandeyes.Agents

	for _, er := range v.([]interface{}) {
		rer := er.(map[string]interface{})
		agent := &thousandeyes.Agent{
			AgentID: rer["agent_id"].(int),
		}
		agents = append(agents, *agent)
	}

	return agents
}

func expandAlertRules(v interface{}) thousandeyes.AlertRules {
	var alertRules thousandeyes.AlertRules

	for _, er := range v.([]interface{}) {
		rer := er.(map[string]interface{})
		alertRule := &thousandeyes.AlertRule{
			RuleID: rer["rule_id"].(int),
		}
		alertRules = append(alertRules, *alertRule)
	}

	return alertRules
}

func expandBGPMonitors(v interface{}) thousandeyes.BGPMonitors {
	var bgpMonitors thousandeyes.BGPMonitors

	for _, er := range v.([]interface{}) {
		rer := er.(map[string]interface{})
		bgpMonitor := &thousandeyes.BGPMonitor{
			MonitorID: rer["monitor_id"].(int),
		}
		bgpMonitors = append(bgpMonitors, *bgpMonitor)
	}

	return bgpMonitors
}

func expandDNSServers(v interface{}) []thousandeyes.Server {
	var dnsServers []thousandeyes.Server

	for _, er := range v.([]interface{}) {
		rer := er.(map[string]interface{})
		targetDNSServer := &thousandeyes.Server{
			ServerName: rer["server_name"].(string),
		}
		dnsServers = append(dnsServers, *targetDNSServer)
	}

	return dnsServers
}

func unpackSIPAuthData(i interface{}) thousandeyes.SIPAuthData {
	var m = i.(map[string]interface{})
	var sipAuthData = thousandeyes.SIPAuthData{}

	for k, v := range m {
		if k == "auth_user" {
			sipAuthData.AuthUser = v.(string)
		}
		if k == "password" {
			sipAuthData.Password = v.(string)
		}
		if k == "port" {
			port, err := strconv.Atoi(v.(string))
			if err == nil {
				sipAuthData.Port = port
			}
		}
		if k == "protocol" {
			sipAuthData.Protocol = v.(string)
		}
		if k == "sip_proxy" {
			sipAuthData.SipProxy = v.(string)
		}
		if k == "sip_registrar" {
			sipAuthData.SipRegistrar = v.(string)
		}
		if k == "user" {
			sipAuthData.User = v.(string)
		}
	}

	return sipAuthData
}

// ResourceBuildStruct builds a struct for a given test type
func ResourceBuildStruct(d *schema.ResourceData, referenceStruct interface{}) interface{} {
	v := reflect.ValueOf(referenceStruct)
	t := reflect.TypeOf(referenceStruct)
	for i := 0; i < v.NumField(); i++ {
		//fieldName := t.Field(i).Name
		tag := GetJSONKey(t.Field(i))
		tfName := CamelCaseToUnderscore(tag)
		val, ok := d.GetOk(tfName)
		if ok {
			newVal := reflect.ValueOf(FillValue(val, t))
			//setField := reflect.ValueOf(&referenceStruct).Elem().Field(i)
			setElem := reflect.ValueOf(&referenceStruct).Elem()
			setField := setElem.Field(i)
			setField.Set(newVal)
		}
	}

	return referenceStruct
}

// ResourceRead sets values for a schema.ResourceData object from a struct
func ResourceRead(d *schema.ResourceData, referenceStruct interface{}) interface{} {
	v := reflect.ValueOf(referenceStruct)
	t := reflect.TypeOf(referenceStruct)
	for i := 0; i < v.NumField(); i++ {
		//fieldName := t.Field(i).Name
		tag := GetJSONKey(t.Field(i))
		tfName := CamelCaseToUnderscore(tag)
		d.Set(tfName, v.Field(i))
	}

	return nil
}

// ResourceUpdate updates a struct's values if changes for those values are
// found in a provided schema.ResourceData object.
func ResourceUpdate(d *schema.ResourceData, referenceStruct interface{}) interface{} {
	//  d.Partial(true)
	//if d.HasChange("name") {
	//	update.TestName = d.Get("name").(string)
	//for field in struct
	d.Partial(true)
	v := reflect.ValueOf(referenceStruct)
	t := reflect.TypeOf(referenceStruct)
	for i := 0; i < v.NumField(); i++ {
		//fieldName := t.Field(i).Name
		tag := GetJSONKey(t.Field(i))
		tfName := CamelCaseToUnderscore(tag)
		if d.HasChange(tfName) {
			newVal := FillValue(d.Get(tfName), v.Field(i))
			v.Field(i).Elem().Set(reflect.ValueOf(newVal))
		}
	}
	d.Partial(false)
	return nil
}

// ResourceSchemaBuild creates a map of schemas based on the fields
// of the provided type.
func ResourceSchemaBuild(referenceStruct interface{}) map[string]*schema.Schema {
	newSchema := map[string]*schema.Schema{}
	v := reflect.ValueOf(referenceStruct)
	t := reflect.TypeOf(referenceStruct)

	for i := 0; i < v.NumField(); i++ {
		//fieldName := t.Field(i).Name
		tag := GetJSONKey(t.Field(i))
		tfName := CamelCaseToUnderscore(tag)
		if val, ok := schemas[tfName]; ok {
			newSchema[tfName] = val
		}
	}
	log.Printf("[INFO] Schema Built")
	return newSchema
}

// FillValue translats a value from the Terraform provider framework and
// translates it to the correct type, based on the type of the target parameter.
func FillValue(source interface{}, target interface{}) interface{} {
	// We determine how to interpret the supplied value based on
	// the type of the target argument.
	switch reflect.ValueOf(target).Kind() {
	case reflect.Slice:
		// When the target is a slice, we create a new slice of the same type,
		// then recurse with the value of each element.
		v := reflect.ValueOf(target)
		t := reflect.TypeOf(target)
		st := reflect.TypeOf(t.Elem())
		newSlice := reflect.Zero(t).Elem()
		for i := 0; i < v.NumField(); i++ {
			newVal := reflect.ValueOf(FillValue(v.Field(i), st))
			//switch st.Kind() {
			//case reflect.Struct:
			//newStruct := reflect.Zero(st)
			//newStructType := reflect.TypeOf(newStruct)
			//for j := 0; j < newStruct.NumField(); j++ {
			//tag := GetJSONKey(newStructType.Field(j))
			//tfName := CamelCaseToUnderscore(tag)
			//rer := value.(map[string]interface{})
			//if sv, ok := rer[tfName]; ok {
			//newStruct.Field(i).Set(sv)
			//}
			//}
			newSlice = reflect.Append(newSlice, newVal)
			//default:
			//newSlice = reflect.Append(newSlice, v)
			//}
		}
		return newSlice
	case reflect.Struct:
		// When the target is a struct, we assume that the source is a map
		// containing corresponding values for the struct's fields.
		v := reflect.ValueOf(target)
		//t := reflect.TypeOf(target)
		for i := 0; i < v.NumField(); i++ {
			newVal := FillValue(source, v.Field(i))
			v.Field(i).Elem().Set(reflect.ValueOf(newVal))
		}
		return v
	case reflect.Int:
		// Int values come to us as strings.
		i, _ := strconv.Atoi(source.(string))
		return i
	default:
		// If we haven't matched one of the above cases, then there
		// is likely no reason to translate.
		return source
	}

	// Or if the above is too tricky...
	//switch t := target.(type) {
	//case []thousandeyes.Agents:
	//return expandAgents(value)
	//case []thousandeyes.AlertRules:
	//return expandAlertRules(value)
	//case []thousandeyes.BGPMonitor:
	//return expandBGPMonitors(value)
	//case []thousandeyes.DNSServer:
	//return expandDNSServers(value)
	//case thousandeyes.SIPAuthData:
	//return unpackSIPAuthData(value)
	//default:
	//return v
	//}
}

// UnderscoreToLowerCamelCase translates from words separated by
// underscores to camel case with initial lowercase.
// ie, a_string would become aString
func UnderscoreToLowerCamelCase(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)
	firstChar := string(s[0])
	s = strings.Replace(s, firstChar, strings.ToLower(firstChar), 1)
	return s
}

// CamelCaseToUnderscore translates from camel case (with any leading case)
// to underscore separated words.
// ie, either aString and AString would become a_string
// Special exception for testName, which becomes "name" to preserve
// pre-existing functionality.
func CamelCaseToUnderscore(s string) string {
	var out []rune
	for i, r := range []rune(s) {
		if unicode.IsUpper(r) {
			if i != 0 {
				out = append(out, []rune("_")[0])
			}
			out = append(out, unicode.ToLower(r))
		} else {
			out = append(out, r)
		}
	}

	outString := string(out)
	if outString == "test_name" {
		outString = "name"
	}
	return outString
}

// GetJSONKey returns the JSON object key for the struct which is represented
// by the passed reflect.StructField instance.
func GetJSONKey(v reflect.StructField) string {
	s := v.Tag.Get("json")
	return strings.Split(s, ",")[0]
}
