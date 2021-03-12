package gen

//go:generate mockgen -package twilio_mock -destination mocktwilio/mock/sms/sms.gen.go github.com/suliar/GFSender/quotes TwilioClient
