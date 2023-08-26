package constant

const (
	SERVICE_NAME    = "VDI-Meter-API"
	SERVICE_VERSION = "1.0.3"

	KEY_PASSWORD = "V1si0n3t*1!"

	TABLE_REASON_NAME         = "problem"
	TABLE_STATUS_NAME         = "status"
	TABLE_STATUS_NAME_PRELOAD = "Status"

	TABLE_UNIT_NAME                = "unit"
	TABLE_UNIT_NAME_PRELOAD        = "Unit"
	TABLE_TASKLIST_NAME            = "tasklist"
	TABLE_TASKLIST_LOG_IMPORT_NAME = "tasklist_log_import"
	TABLE_PROBLEM                  = "problem"
	TABLE_TASKLIST_NAME_PRELOAD    = "Tasklist"
	TABLE_TASKLIST_METER           = "tasklist_meter"
	TABLE_TASKLIST_METER_PRELOAD   = "TasklistMeter"
	TABLE_FEATURE_NAME             = "feature"
	TABLE_FEATURE_NAME_PRELOAD     = "Feature"

	TABLE_MASTER_SUPPLIER  = "MasterSupplier"
	TABLE_MASTER_CUSTOMER  = "MasterCustomer"
	TABLE_MASTER_PRODUCT   = "MasterProduct"
	TABLE_MASTER_WAREHOUSE = "MasterWarehouse"

	TABLE_TRX_PENERIMAAN_BARANG_HEADER  = "TrxPenerimaanBarangHeader"
	TABLE_TRX_PENERIMAAN_BARANG_DETAIL  = "TrxPenerimaanBarangDetail"
	TABLE_TRX_PENGELUARAN_BARANG_HEADER = "TrxPengeluaranBarangHeader"
	TABLE_TRX_PENGELUARAN_BARANG_DETAIL = "TrxPengeluaranBarangDetail"

	TABLE_ROLE_NAME            = "role"
	TABLE_ROLE_PERMISSION_NAME = "role_permission"

	TABLE_USERS                        = "users"
	TABLE_USER_LOG_NAME                = "user_log"
	TABLE_USER_ACTIVITY_LOG_NAME       = "user_activity_log"
	TABLE_USER_NAME_PRELOAD            = "User"
	TABLE_USER_CREDENTIAL_NAME         = "user_credential"
	TABLE_USER_CREDENTIAL_NAME_PRELOAD = "UserCredential"
	TABLE_USER_ROLE_NAME               = "user_role"

	TABLE_OTP_NAME = "otp"

	TABLE_CONFIG_NAME = "config"

	OPERATION_SQL_INSERT = "insert"
	OPERATION_SQL_UPDATE = "update"
	OPERATION_SQL_DELETE = "delete"

	FEATURE_ROLE_ID = 7
	FEATURE_USER_ID = 6

	ACTION_UPDATE = "UPDATE"
	ACTION_DELETE = "DELETE"
	ACTION_CREATE = "CREATE"
	ACTION_IMPORT = "IMPORT"

	IMAGE_PNG_EXTENSION = "png"

	CONTENT_TYPE_IMAGE_PNG        = "image/png"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"

	FORMAT_TIME_SYNC                 = "2006-01-02 15:04:05 -0700"
	FORMAT_TIME_SYNCWITHOUT_TIMEZONE = "2006-01-02 15:04:05"

	DEFAULT_QUERY_SOFT_DELETE = "deleted_at IS NULL"

	ENV_MINIO_DOMAIN = "MINIO_DOMAIN"
	MINIO_BUCKET     = "vdi-meter"
)
