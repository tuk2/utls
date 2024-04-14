package dicttls

// source: https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#tls-parameters-8
// last updated: March 2023

const (
	SupportedGroups_x25519                       	uint16 = 1
	SupportedGroups_secp256r1                       uint16 = 2
	SupportedGroups_secp384r1                       uint16 = 3
	SupportedGroups_secp521r1                       uint16 = 4
)

var DictSupportedGroupsValueIndexed = map[uint16]string{
	1:    "x25519",
	2:    "secp256r1",
	3:    "secp384r1",
	4:    "secp521r1",
}

var DictSupportedGroupsNameIndexed = map[string]uint16{
	"x25519":                       1,
	"secp256r1":                       2,
	"secp384r1":                       3,
	"secp521r1":                       4,
}
