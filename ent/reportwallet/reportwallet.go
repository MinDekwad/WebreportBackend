// Code generated by entc, DO NOT EDIT.

package reportwallet

const (
	// Label holds the string label denoting the reportwallet type in the database.
	Label = "report_wallet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldWalletid holds the string denoting the walletid field in the database.
	FieldWalletid = "WalletID"
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
	// FieldDateTime holds the string denoting the datetime field in the database.
	FieldDateTime = "DateTime"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "Balance"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "Email"
	// FieldIsForgetPin holds the string denoting the isforgetpin field in the database.
	FieldIsForgetPin = "IsForgetPin"
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
	// FieldRegisterDateTime holds the string denoting the registerdatetime field in the database.
	FieldRegisterDateTime = "RegisterDateTime"
	// FieldFileimportID holds the string denoting the fileimportid field in the database.
	FieldFileimportID = "FileimportID"

	// Table holds the table name of the reportwallet in the database.
	Table = "wallet_daily"
)

// Columns holds all SQL columns for reportwallet fields.
var Columns = []string{
	FieldID,
	FieldWalletid,
	FieldWalletTypeName,
	FieldWalletPhoneno,
	FieldWalletName,
	FieldCitizenId,
	FieldStatus,
	FieldDateTime,
	FieldBalance,
	FieldEmail,
	FieldIsForgetPin,
	FieldATMCard,
	FieldAccountNo,
	FieldAddressDetail,
	FieldStreet,
	FieldDistrict,
	FieldSubDistrict,
	FieldProvince,
	FieldPostalCode,
	FieldRegisterDateTime,
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
	// EmailValidator is a validator for the "Email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// IsForgetPinValidator is a validator for the "IsForgetPin" field. It is called by the builders before save.
	IsForgetPinValidator func(string) error
	// ATMCardValidator is a validator for the "ATMCard" field. It is called by the builders before save.
	ATMCardValidator func(string) error
	// AccountNoValidator is a validator for the "AccountNo" field. It is called by the builders before save.
	AccountNoValidator func(string) error
	// StreetValidator is a validator for the "Street" field. It is called by the builders before save.
	StreetValidator func(string) error
	// DistrictValidator is a validator for the "District" field. It is called by the builders before save.
	DistrictValidator func(string) error
	// SubDistrictValidator is a validator for the "SubDistrict" field. It is called by the builders before save.
	SubDistrictValidator func(string) error
	// ProvinceValidator is a validator for the "Province" field. It is called by the builders before save.
	ProvinceValidator func(string) error
	// PostalCodeValidator is a validator for the "PostalCode" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
)
