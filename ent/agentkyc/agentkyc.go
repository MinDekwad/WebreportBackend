// Code generated by entc, DO NOT EDIT.

package agentkyc

const (
	// Label holds the string label denoting the agentkyc type in the database.
	Label = "agentkyc"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldKYCDate holds the string denoting the kycdate field in the database.
	FieldKYCDate = "KYCDate"
	// FieldKYCTime holds the string denoting the kyctime field in the database.
	FieldKYCTime = "KYCTime"
	// FieldAgentID holds the string denoting the agentid field in the database.
	FieldAgentID = "AgentID"
	// FieldAgentemail holds the string denoting the agentemail field in the database.
	FieldAgentemail = "Agentemail"
	// FieldAgentNameLastname holds the string denoting the agentnamelastname field in the database.
	FieldAgentNameLastname = "AgentNameLastname"
	// FieldKYCStatus holds the string denoting the kycstatus field in the database.
	FieldKYCStatus = "KYCStatus"
	// FieldConsumerwalletid holds the string denoting the consumerwalletid field in the database.
	FieldConsumerwalletid = "Consumerwalletid"
	// FieldKYCRespond holds the string denoting the kycrespond field in the database.
	FieldKYCRespond = "KYCRespond"
	// FieldDOPARespond holds the string denoting the doparespond field in the database.
	FieldDOPARespond = "DOPARespond"
	// FieldAgentType holds the string denoting the agenttype field in the database.
	FieldAgentType = "AgentType"
	// FieldFileimportID holds the string denoting the fileimportid field in the database.
	FieldFileimportID = "FileimportID"

	// Table holds the table name of the agentkyc in the database.
	Table = "agent_kyc"
)

// Columns holds all SQL columns for agentkyc fields.
var Columns = []string{
	FieldID,
	FieldKYCDate,
	FieldKYCTime,
	FieldAgentID,
	FieldAgentemail,
	FieldAgentNameLastname,
	FieldKYCStatus,
	FieldConsumerwalletid,
	FieldKYCRespond,
	FieldDOPARespond,
	FieldAgentType,
	FieldFileimportID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// KYCDateValidator is a validator for the "KYCDate" field. It is called by the builders before save.
	KYCDateValidator func(string) error
	// KYCTimeValidator is a validator for the "KYCTime" field. It is called by the builders before save.
	KYCTimeValidator func(string) error
	// AgentIDValidator is a validator for the "AgentID" field. It is called by the builders before save.
	AgentIDValidator func(string) error
	// AgentemailValidator is a validator for the "Agentemail" field. It is called by the builders before save.
	AgentemailValidator func(string) error
	// AgentNameLastnameValidator is a validator for the "AgentNameLastname" field. It is called by the builders before save.
	AgentNameLastnameValidator func(string) error
	// KYCStatusValidator is a validator for the "KYCStatus" field. It is called by the builders before save.
	KYCStatusValidator func(string) error
	// ConsumerwalletidValidator is a validator for the "Consumerwalletid" field. It is called by the builders before save.
	ConsumerwalletidValidator func(string) error
	// KYCRespondValidator is a validator for the "KYCRespond" field. It is called by the builders before save.
	KYCRespondValidator func(string) error
	// DOPARespondValidator is a validator for the "DOPARespond" field. It is called by the builders before save.
	DOPARespondValidator func(string) error
	// AgentTypeValidator is a validator for the "AgentType" field. It is called by the builders before save.
	AgentTypeValidator func(string) error
)