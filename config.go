package main

var Config = struct {
	OwnerNumber   string
	BotName       string
	OTPChannelIDs []string
	OTPApiURLs    []string
	Interval      int
}{
	OwnerNumber: "923273788442",
	BotName:     "Zero OTP Monitor",
	OTPChannelIDs: []string{
		"120363423562861659@newsletter",
		"120363407230990898@newsletter",
	},
	OTPApiURLs: []string{
		"https://ali-api-proo.up.railway.app/api/np?type=sms",
		"https://ali-api-proo.up.railway.app/api/msi?type=sms",
		"https://ali-api-proo.up.railway.app/api/mat?type=sms",
		"https://ali-api-proo.up.railway.app/api/ts?type=sms",
		"https://ali-api-proo.up.railway.app/api/ch?type=sms",
		"https://ali-api-proo.up.railway.app/api/gen?type=sms",
		"https://ali-api-proo.up.railway.app/api/ivs?type=sms",
	},
	Interval: 2, // ✅ 3 sec - faster than before
}
