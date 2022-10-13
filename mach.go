// match matchs common regular expression pattrens
package match

import "regexp"

// IsRegular expression patterns
const (
	DatePattern           = `(?i)(?:[0-3]?\d(?:st|nd|rd|th)?\s+(?:of\s+)?(?:jan\.?|january|feb\.?|february|mar\.?|march|apr\.?|april|may|jun\.?|june|jul\.?|july|aug\.?|august|sep\.?|september|oct\.?|october|nov\.?|november|dec\.?|december)|(?:jan\.?|january|feb\.?|february|mar\.?|march|apr\.?|april|may|jun\.?|june|jul\.?|july|aug\.?|august|sep\.?|september|oct\.?|october|nov\.?|november|dec\.?|december)\s+[0-3]?\d(?:st|nd|rd|th)?)(?:\,)?\s*(?:\d{4})?|[0-3]?\d[-\./][0-3]?\d[-\./]\d{2,4}`
	TimePattern           = `(?i)\d{1,2}:\d{2} ?(?:[ap]\.?m\.?)?|\d[ap]\.?m\.?`
	PhonePattern          = `(?:(?:\+?\d{1,3}[-.\s*]?)?(?:\(?\d{3}\)?[-.\s*]?)?\d{3}[-.\s*]?\d{4,6})|(?:(?:(?:\(\+?\d{2}\))|(?:\+?\d{2}))\s*\d{2}\s*\d{3}\s*\d{4})`
	PhonesWithExtsPattern = `(?i)(?:(?:\+?1\s*(?:[.-]\s*)?)?(?:\(\s*(?:[2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9])\s*\)|(?:[2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9]))\s*(?:[.-]\s*)?)?(?:[2-9]1[02-9]|[2-9][02-9]1|[2-9][02-9]{2})\s*(?:[.-]\s*)?(?:[0-9]{4})(?:\s*(?:#|x\.?|ext\.?|extension)\s*(?:\d+)?)`
	LinkPattern           = `(?:(?:https?:\/\/)?(?:[a-z0-9.\-]+|www|[a-z0-9.\-])[.](?:[^\s()<>]+|\((?:[^\s()<>]+|(?:\([^\s()<>]+\)))*\))+(?:\((?:[^\s()<>]+|(?:\([^\s()<>]+\)))*\)|[^\s!()\[\]{};:\'".,<>?]))`
	EmailPattern          = `(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`
	IPv4Pattern           = `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`
	IPv6Pattern           = `(?:(?:(?:[0-9A-Fa-f]{1,4}:){7}(?:[0-9A-Fa-f]{1,4}|:))|(?:(?:[0-9A-Fa-f]{1,4}:){6}(?::[0-9A-Fa-f]{1,4}|(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(?:(?:[0-9A-Fa-f]{1,4}:){5}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,2})|:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(?:(?:[0-9A-Fa-f]{1,4}:){4}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,3})|(?:(?::[0-9A-Fa-f]{1,4})?:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){3}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,4})|(?:(?::[0-9A-Fa-f]{1,4}){0,2}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){2}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,5})|(?:(?::[0-9A-Fa-f]{1,4}){0,3}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?:(?:[0-9A-Fa-f]{1,4}:){1}(?:(?:(?::[0-9A-Fa-f]{1,4}){1,6})|(?:(?::[0-9A-Fa-f]{1,4}){0,4}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(?::(?:(?:(?::[0-9A-Fa-f]{1,4}){1,7})|(?:(?::[0-9A-Fa-f]{1,4}){0,5}:(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(?:%.+)?\s*`
	IPPattern             = IPv4Pattern + `|` + IPv6Pattern
	NotKnownPortPattern   = `6[0-5]{2}[0-3][0-5]|[1-5][\d]{4}|[2-9][\d]{3}|1[1-9][\d]{2}|10[3-9][\d]|102[4-9]`
	PricePattern          = `[$]\s?[+-]?[0-9]{1,3}(?:(?:,?[0-9]{3}))*(?:\.[0-9]{1,2})?`
	HexColorPattern       = `(?:#?([0-9a-fA-F]{6}|[0-9a-fA-F]{3}))`
	CreditCardPattern     = `(?:(?:(?:\d{4}[- ]?){3}\d{4}|\d{15,16}))`
	VISACreditCardPattern = `4\d{3}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`
	MCCreditCardPattern   = `5[1-5]\d{2}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`
	BtcAddressPattern     = `[13][a-km-zA-HJ-NP-Z1-9]{25,34}`
	StreetAddressPattern  = `\d{1,4} [\w\s]{1,20}(?:street|st|avenue|ave|road|rd|highway|hwy|square|sq|trail|trl|drive|dr|court|ct|park|parkway|pkwy|circle|cir|boulevard|blvd)\W?`
	ZipCodePattern        = `\b\d{5}(?:[-\s]\d{4})?\b`
	PoBoxPattern          = `(?i)P\.? ?O\.? Box \d+`
	SSNPattern            = `(?:\d{3}-\d{2}-\d{4})`
	MD5HexPattern         = `[0-9a-fA-F]{32}`
	SHA1HexPattern        = `[0-9a-fA-F]{40}`
	SHA256HexPattern      = `[0-9a-fA-F]{64}`
	GUIDPattern           = `[0-9a-fA-F]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}`
	ISBN13Pattern         = `(?:[\d]-?){12}[\dxX]`
	ISBN10Pattern         = `(?:[\d]-?){9}[\dxX]`
	MACAddressPattern     = `(([a-fA-F0-9]{2}[:-]){5}([a-fA-F0-9]{2}))`
	IBANPattern           = `[A-Z]{2}\d{2}[A-Z0-9]{4}\d{7}([A-Z\d]?){0,16}`
	GitRepoPattern        = `((git|ssh|http(s)?)|(git@[\w\.]+))(:(\/\/)?)([\w\.@\:/\-~]+)(\.git)(\/)?`
)

// IsCompiled regular expressions
var (
	DateRegex           = regexp.MustCompile(DatePattern)
	TimeRegex           = regexp.MustCompile(TimePattern)
	PhoneRegex          = regexp.MustCompile(PhonePattern)
	PhonesWithExtsRegex = regexp.MustCompile(PhonesWithExtsPattern)
	LinkRegex           = regexp.MustCompile(LinkPattern)
	EmailRegex          = regexp.MustCompile(EmailPattern)
	IPv4Regex           = regexp.MustCompile(IPv4Pattern)
	IPv6Regex           = regexp.MustCompile(IPv6Pattern)
	IPRegex             = regexp.MustCompile(IPPattern)
	NotKnownPortRegex   = regexp.MustCompile(NotKnownPortPattern)
	PriceRegex          = regexp.MustCompile(PricePattern)
	HexColorRegex       = regexp.MustCompile(HexColorPattern)
	CreditCardRegex     = regexp.MustCompile(CreditCardPattern)
	BtcAddressRegex     = regexp.MustCompile(BtcAddressPattern)
	StreetAddressRegex  = regexp.MustCompile(StreetAddressPattern)
	ZipCodeRegex        = regexp.MustCompile(ZipCodePattern)
	PoBoxRegex          = regexp.MustCompile(PoBoxPattern)
	SSNRegex            = regexp.MustCompile(SSNPattern)
	MD5HexRegex         = regexp.MustCompile(MD5HexPattern)
	SHA1HexRegex        = regexp.MustCompile(SHA1HexPattern)
	SHA256HexRegex      = regexp.MustCompile(SHA256HexPattern)
	GUIDRegex           = regexp.MustCompile(GUIDPattern)
	ISBN13Regex         = regexp.MustCompile(ISBN13Pattern)
	ISBN10Regex         = regexp.MustCompile(ISBN10Pattern)
	VISACreditCardRegex = regexp.MustCompile(VISACreditCardPattern)
	MCCreditCardRegex   = regexp.MustCompile(MCCreditCardPattern)
	MACAddressRegex     = regexp.MustCompile(MACAddressPattern)
	IBANRegex           = regexp.MustCompile(IBANPattern)
	GitRepoRegex        = regexp.MustCompile(GitRepoPattern)
)

func match(text string, regex *regexp.Regexp) bool {
	// MatchString is faster than Mach ([]byte) in most cases.
	return regex.MatchString(text)
}

// IsDate finds all date strings
func IsDate(text string) bool {
	return match(text, DateRegex)
}

// IsTime finds all time strings
func IsTime(text string) bool {
	return match(text, TimeRegex)
}

// IsPhones finds all phone numbers
func IsPhone(text string) bool {
	return match(text, PhoneRegex)
}

// IsPhonesWithExts finds all phone numbers with ext
func IsPhoneWithExt(text string) bool {
	return match(text, PhonesWithExtsRegex)
}

// IsLinks finds all link strings
func IsLink(text string) bool {
	return match(text, LinkRegex)
}

// IsEmails finds all email strings
func IsEmail(text string) bool {
	return match(text, EmailRegex)
}

// IsIPv4s finds all IPv4 addresses
func IsIPv4(text string) bool {
	return match(text, IPv4Regex)
}

// IsIPv6s finds all IPv6 addresses
func IsIPv6(text string) bool {
	return match(text, IPv6Regex)
}

// IsIPs finds all IP addresses (both IPv4 and IPv6)
func IsIP(text string) bool {
	return match(text, IPRegex)
}

// IsNotKnownPorts finds all not-known port numbers
func IsNotKnownPort(text string) bool {
	return match(text, NotKnownPortRegex)
}

// IsPrices finds all price strings
func IsPrice(text string) bool {
	return match(text, PriceRegex)
}

// IsHexColors finds all hex color values
func IsHexColor(text string) bool {
	return match(text, HexColorRegex)
}

// IsCreditCards finds all credit card numbers
func IsCreditCard(text string) bool {
	return match(text, CreditCardRegex)
}

// IsBtcAddresses finds all bitcoin addresses
func IsBtcAddresse(text string) bool {
	return match(text, BtcAddressRegex)
}

// IsStreetAddresses finds all street addresses
func IsStreetAddresse(text string) bool {
	return match(text, StreetAddressRegex)
}

// IsZipCodes finds all zip codes
func IsZipCode(text string) bool {
	return match(text, ZipCodeRegex)
}

// IsPoBoxes finds all po-box strings
func IsPoBoxe(text string) bool {
	return match(text, PoBoxRegex)
}

// IsSSNs finds all SSN strings
func IsSSN(text string) bool {
	return match(text, SSNRegex)
}

// IsMD5Hexes finds all MD5 hex strings
func IsMD5Hexe(text string) bool {
	return match(text, MD5HexRegex)
}

// IsSHA1Hexes finds all SHA1 hex strings
func IsSHA1Hexe(text string) bool {
	return match(text, SHA1HexRegex)
}

// IsSHA256Hexes finds all SHA256 hex strings
func IsSHA256Hexe(text string) bool {
	return match(text, SHA256HexRegex)
}

// IsGUIDs finds all GUID strings
func IsGUID(text string) bool {
	return match(text, GUIDRegex)
}

// IsISBN13s finds all ISBN13 strings
func IsISBN13(text string) bool {
	return match(text, ISBN13Regex)
}

// IsISBN10s finds all ISBN10 strings
func IsISBN10(text string) bool {
	return match(text, ISBN10Regex)
}

// IsVISACreditCards finds all VISA credit card numbers
func IsVISACreditCard(text string) bool {
	return match(text, VISACreditCardRegex)
}

// IsMCCreditCards finds all MasterCard credit card numbers
func IsMCCreditCard(text string) bool {
	return match(text, MCCreditCardRegex)
}

// IsMACAddresses finds all MAC addresses
func IsMACAddresse(text string) bool {
	return match(text, MACAddressRegex)
}

// IsIBANs finds all IBAN strings
func IsIBAN(text string) bool {
	return match(text, IBANRegex)
}

// IsGitRepos finds all git repository addresses which have protocol prefix
func IsGitRepo(text string) bool {
	return match(text, GitRepoRegex)
}
