// Code generated by entc, DO NOT EDIT.

package billingstatus

const (
	// Label holds the string label denoting the billingstatus type in the database.
	Label = "billingstatus"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBillingstatusname holds the string denoting the billingstatusname field in the database.
	FieldBillingstatusname = "billingstatusname"

	// EdgeBillingstatuss holds the string denoting the billingstatuss edge name in mutations.
	EdgeBillingstatuss = "billingstatuss"

	// Table holds the table name of the billingstatus in the database.
	Table = "billingstatuses"
	// BillingstatussTable is the table the holds the billingstatuss relation/edge.
	BillingstatussTable = "bills"
	// BillingstatussInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillingstatussInverseTable = "bills"
	// BillingstatussColumn is the table column denoting the billingstatuss relation/edge.
	BillingstatussColumn = "billingstatus_id"
)

// Columns holds all SQL columns for billingstatus fields.
var Columns = []string{
	FieldID,
	FieldBillingstatusname,
}
