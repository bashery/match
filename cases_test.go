package match

var trueEmails = []string{
	"adam@email.com",
	"joha@gmail.com",
	"koko_helo@hotmail.com",
	"ahmed@bashery.net",
}

var falseEmails = []string{
	"adam#gmail.com",
	"joya_emmail.net",
	"jawad@hotmail@org",
	"55:59:59",
	"01:66:59",
	"23:23:oo",
}

var trueTime = []string{
	"23:59:59",
	"00:00:00",
	"12:12:12",
	"00:56:56",
	"12:13:12",
	"12:12:01",
	"09:45",
	"9:45",
	"23:45",
	"9:00am",
	"9am",
	"9:00 A.M.",
	"9:00 pm",
}

var falseTime = []string{
	"55:59:59",
	"01:66:59",
	"23:23:oo",
	"01:66:59",
	"01:66:59",
	"33:66:59",
}

///////////////////////

var trueLinks = []string{
	"www.google.com",
	"http://www.google.com",
	"www.google.com/?query=dog",
	"sub.example.com",
	"http://www.google.com/%&#/?q=dog",
	"google.com",
}

var falseLinks = []string{
	"www#google.com",
	"http://www.google@com",
	"google#com/?query=dog",
	"##.e.x.ample.com",
	"hp://ww.co.com/%&#/?q=dog",
	"google.000",
}

// Dates
var trueDates = []string{
	"3-23-17",
	"3.23.17",
	"03.23.17",
	"March 23th, 2017",
	"Mar 23th 2017",
	"Mar. 23th, 2017",
	"23 Mar 2017",
}

var tPhons = []string{
	"12345678900",
	"1234567890",
	"+1 234 567 8900",
	"234-567-8900",
	"1-234-567-8900",
	"1.234.567.8900",
	"5678900",
	"567-8900",
	"(003) 555-1212",
	"+41 22 730 5989",
	"+442345678900",
}

var tPhonesWithExts = []string{
	"(523)222-8888 ext 527",
	"(523)222-8888x623",
	"(523)222-8888 x623",
	"(523)222-8888 x 623",
	"(523)222-8888EXT623",
	"523-222-8888EXT623",
	"(523) 222-8888 x 623",
}

var tLinks = []string{
	"www.google.com",
	"http://www.google.com",
	"www.google.com/?query=dog",
	"sub.example.com",
	"http://www.google.com/%&#/?q=dog",
	"google.com",
}

var tEmails = []string{
	"john.smith@gmail.com",
	"john_smith@gmail.com",
	"john@example.net",
	"John@example.net",
}

var IPv4s = []string{
	"127.0.0.1",
	"192.168.1.1",
	"8.8.8.8",
	"192.30.253.113",
	"216.58.194.46",
}
var IPv6s = []string{
	"fe80:0000:0000:0000:0204:61ff:fe9d:f156",
	"fe80:0:0:0:204:61ff:fe9d:f156",
	"fe80::204:61ff:fe9d:f156",
	"fe80:0000:0000:0000:0204:61ff:254.157.241.86",
	"fe80:0:0:0:0204:61ff:254.157.241.86",
	"::1",
}

var IPs = []string{
	"127.0.0.1",
	"192.168.1.1",
	"8.8.8.8",
	"192.30.253.113",
	"216.58.194.46",
	"fe80:0000:0000:0000:0204:61ff:fe9d:f156",
	"fe80:0:0:0:204:61ff:fe9d:f156",
	"fe80::204:61ff:fe9d:f156",
	"fe80:0000:0000:0000:0204:61ff:254.157.241.86",
	"fe80:0:0:0:0204:61ff:254.157.241.86",
	"::1",
}

var trueNotKnownPorts = []string{
	"1024",
	"2121",
	"8080",
	"12345",
	"55555",
	"65535",
}

var faseNotKnownPorts = []string{
	"21",
	"80",
	"1023",
	"65536",
}

var tPrices = []string{
	"$1.23",
	"$1",
	"$1,000",
	"$10,000.00",
}

var failingPrices = []string{
	"$1,10,0",
	"$100.000",
}

var trueHexColors = []string{
	"#fff",
	"#123",
	"#4e32ff",
}

var failingHexColors = []string{
	"#zzz",
}

var trueCreditCards = []string{
	"0000-0000-0000-0000",
	"0123456789012345",
	"0000 0000 0000 0000",
	"012345678901234",
}

var trueBtcAddresses = []string{
	"1LgqButDNV2rVHe9DATt6WqD8tKZEKvaK2",
	"19P6EYhu6kZzRy9Au4wRRZVE8RemrxPbZP",
	"1bones8KbQge9euDn523z5wVhwkTP3uc1",
	"1Bow5EMqtDGV5n5xZVgdpRPJiiDK6XSjiC",
}

var falseBtcAddresses = []string{
	"2LgqButDNV2rVHe9DATt6WqD8tKZEKvaK2",
	"19Ry9Au4wRRZVE8RemrxPbZP",
	"1bones8KbQge9euDn523z5wVhwkTP3uc12939",
	"1Bow5EMqtDGV5n5xZVgdpR",
}

var trueStreetAddresses = []string{
	"101 main st.",
	"504 parkwood drive",
	"3 elm boulevard",
	"500 elm street ",
}

var falseStreetAddresses = []string{
	"101 main straight",
}

var trueZipCodes = []string{
	"02540",
	"02540-4119",
}

var falseZipCodes = []string{
	"101 main straight",
	"123456",
}

var truePoBoxes = []string{
	"PO Box 123456",
	"p.o. box 234234",
}

var falsePoBoxes = []string{
	"101 main straight",
}

var trueSSNs = []string{
	"000-00-0000",
	"111-11-1111",
	"222-22-2222",
	"123-45-6789",
}

var trueMD5Hexes = []string{
	"b5ab01fad5a008d436f76aafc896f9c6",
	"00000000000000000000000000000000",
	"fffFFFfFFfFFFfFFFFfFfFfffffFfFFF",
}

var falseMD5Hexes = []string{
	"b5ab01fad5a008d436f76aafc896f9c600000000",
	"",
	"7TS5x1trQs652k4AZ3hJE83YCvJRy0U8",
	"b5ab01fad5a008-436f76aafc896f9c6",
}

var tureSHA1Hexes = []string{
	"b5ab01fad5a008d436f76aafc896f9c6abcd1234",
	"0000000000000000000000000000000000000000",
	"fffFFFfFFfFFFfFFFFfFfFfffffFfFFFffffFFFF",
}

var falseSHA1Hexes = []string{
	"b5ab01fad5a008d436f76aafc896f9c600000000202020202020202020202020",
	"7TS5x1trQs652k4AZ3hJE83YCvJRy0U85x1trQs652k4AZ3hJE83YCvJRy0U8asd",
	"b5ab01fad5a008-436f76aafc896f9c6-436f76aafc896f9c6-436f76aafc896",
	"",
}

var trueSHA256Hexes = []string{
	"3f4146a1d0b5dac26562ff7dc6248573f4e996cf764a0f517318ff398dcfa792",
	"0000000000000000000000000000000000000000000000000000000000000000",
	"fffFFFfFFfFFFfFFFFfFfFfffffFfFFFffffFFFFfffffFFFFFffFFffFFffFFff",
}

var falseSHA256Hexes = []string{
	"3f4146a1d0b5dac26562ff7dc6248573f4e996cf764a0f517318ff398dcfa7920",
	"",
	"e9iLS075z9HAJlUWg2ZpK5hRxjLeSpIqMKJO67c739VYf7Bj7eR1WjOO82IHcXVd",
	"b5ab01fad5a008-436f76aafc896f9c6-436f76aafc896f9c6-436f76aafc896",
}
var trueGUIDs = []string{
	"00000000-0000-0000-0000-000000000000",
	"00000000000000000000000000000000",
	"88a310ed-0ac0-4a3d-b3a2-958fa291d061",
	"27143ecab8a440cda6fb6effcf9b3c75",
}

var falseGUIDs = []string{
	"88a310ed-0ac0_4a3d_b3a2_958fa291d061",
	"88a310ed 0ac0 4a3d b3a2 958fa291d061",
	"",
	"Z8a310ed-0ac0-4a3d-b3a2-958fa291d061",
	"88a310ed-zac0-4a3d-b3a2-958fa291d061",
	"98a310ed-0ac0-za3d-b3a2-958fa291d061",
	"88a310ed-0ac0-4a3d-z3a2-958fa291d061",
	"88a310ed-0ac0-4a3d-b3a2-z58fa291d061",
}

var trueISBN13s = []string{
	"978-3-16-148410-0",
	"978-1-56619-909-4",
	"133-1-12144-909-9",
}

var falseISBN13s = []string{
	"1-56619-909-3",
	"1-33342-100-1",
	"2-33342-362-9",
}

var trueISBN10s = []string{
	"1-56619-909-3",
	"1-33342-100-1",
	"2-33342-362-9",
}

var falseISBN10s = []string{
	"978-3-16-148410-0",
	"978-1-56619-909-4",
	"133-1-12144-909-9",
}

var trueVISACreditCards = []string{
	"4111 1111 1111 1111",
	"4222 2222 2222 2222",
}

var falseVISACreditCards = []string{
	"5500 0000 0000 0004",
	"3400 0000 0000 009",
	"3000 0000 0000 04",
}

var tureMCCreditCards = []string{
	"5500 0000 0000 0004",
	"5500 3334 0000 1234",
}

var falseMCCreditCards = []string{
	"4111 1111 1111 1111",
	"4222 2222 2222 2222",
	"3400 0000 0000 009",
	"3000 0000 0000 04",
}

var trueMACAddresses = []string{
	"f8:2f:a4:fe:76:d2",
	"F8:2F:A4:FE:76:D2",
	"3D-F2-C9-A6-B3-4F",
}

var falseMACAddresses = []string{
	"3D:F2:C9:A6:B3:4G",
	"f0:2f:P4:Be:96:J5",
}

var trueIBANs = []string{
	"FR1420041010050500013M02606",
	"MU17BOMM0101101030300200000MUR",
	"NO9386011117947",
}

var falseIBANs = []string{
	"424220041010050500013M02606",
	"GB29RBOS601613",
}

var trueGitRepos = []string{
	"https://github.com/mingrammer/commonregex.git",
	"git@github.com:mingrammer/commonregex.git",
}

var falseGitRepos = []string{
	"https://github.com/mingrammer/commonregex",
	"test@github.com:mingrammer/commonregex.git",
}
