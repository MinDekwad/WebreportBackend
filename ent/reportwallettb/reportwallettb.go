// Code generated by entc, DO NOT EDIT.

package reportwallettb

const (
	// Label holds the string label denoting the reportwallettb type in the database.
	Label = "reportwallettb"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldWalletid holds the string denoting the walletid field in the database.
	FieldWalletid = "walletid"
	// FieldWalletTypeName holds the string denoting the wallettypename field in the database.
	FieldWalletTypeName = "WalletTypeName"
	// FieldWalletPhoneno holds the string denoting the walletphoneno field in the database.
	FieldWalletPhoneno = "WalletPhoneno"
	// FieldWalletName holds the string denoting the walletname field in the database.
	FieldWalletName = "WalletName"
	// FieldCitizenId holds the string denoting the citizenid field in the database.
	FieldCitizenId = "CitizenId"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "Status"
	// FieldRegisterDate holds the string denoting the registerdate field in the database.
	FieldRegisterDate = "RegisterDate"
	// FieldGroupUser holds the string denoting the groupuser field in the database.
	FieldGroupUser = "GroupUser"
	// FieldUserAgent holds the string denoting the useragent field in the database.
	FieldUserAgent = "UserAgent"
	// FieldKYCDate holds the string denoting the kyc_date field in the database.
	FieldKYCDate = "KYC_Date"
	// FieldATMCard holds the string denoting the atmcard field in the database.
	FieldATMCard = "ATMCard"
	// FieldAccountNo holds the string denoting the accountno field in the database.
	FieldAccountNo = "AccountNo"
	// FieldAddressDetail holds the string denoting the addressdetail field in the database.
	FieldAddressDetail = "AddressDetail"
	// FieldStreet holds the string denoting the street field in the database.
	FieldStreet = "Street"
	// FieldDistrict holds the string denoting the district field in the database.
	FieldDistrict = "District"
	// FieldSubDistrict holds the string denoting the subdistrict field in the database.
	FieldSubDistrict = "SubDistrict"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "Province"
	// FieldPostalCode holds the string denoting the postalcode field in the database.
	FieldPostalCode = "PostalCode"
	// FieldIsKYC holds the string denoting the iskyc field in the database.
	FieldIsKYC = "isKYC"
	// FieldUpdateDate holds the string denoting the updatedate field in the database.
	FieldUpdateDate = "UpdateDate"

	// Table holds the table name of the reportwallettb in the database.
	Table = "report_wallet"
)

// Columns holds all SQL columns for reportwallettb fields.
var Columns = []string{
	FieldID,
	FieldWalletid,
	FieldWalletTypeName,
	FieldWalletPhoneno,
	FieldWalletName,
	FieldCitizenId,
	FieldStatus,
	FieldRegisterDate,
	FieldGroupUser,
	FieldUserAgent,
	FieldKYCDate,
	FieldATMCard,
	FieldAccountNo,
	FieldAddressDetail,
	FieldStreet,
	FieldDistrict,
	FieldSubDistrict,
	FieldProvince,
	FieldPostalCode,
	FieldIsKYC,
	FieldUpdateDate,
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
	// WalletidValidator is a validator for the "walletid" field. It is called by the builders before save.
	WalletidValidator func(string) error
	// WalletTypeNameValidator is a validator for the "WalletTypeName" field. It is called by the builders before save.
	WalletTypeNameValidator func(string) error
	// WalletPhonenoValidator is a validator for the "WalletPhoneno" field. It is called by the builders before save.
	WalletPhonenoValidator func(string) error
	// WalletNameValidator is a validator for the "WalletName" field. It is called by the builders before save.
	WalletNameValidator func(string) error
	// CitizenIdValidator is a validator for the "CitizenId" field. It is called by the builders before save.
	CitizenIdValidator func(string) error
	// StatusValidator is a validator for the "Status" field. It is called by the builders before save.
	StatusValidator func(string) error
	// UserAgentValidator is a validator for the "UserAgent" field. It is called by the builders before save.
	UserAgentValidator func(string) error
	// ATMCardValidator is a validator for the "ATMCard" field. It is called by the builders before save.
	ATMCardValidator func(string) error
	// AccountNoValidator is a validator for the "AccountNo" field. It is called by the builders before save.
	AccountNoValidator func(string) error
	// ProvinceValidator is a validator for the "Province" field. It is called by the builders before save.
	ProvinceValidator func(string) error
	// PostalCodeValidator is a validator for the "PostalCode" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
	// IsKYCValidator is a validator for the "isKYC" field. It is called by the builders before save.
	IsKYCValidator func(string) error
)