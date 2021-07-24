package gen

//go:generate mockgen -package twilio_mock -destination mocktwilio/mock/sms/sms.gen.go github.com/suliar/GFSender/internal/twilio Client
//go:generate mockgen -package bible_mock -destination mockbible/mock/bibl/bible.gen.go github.com/suliar/GFSender/internal/bible Quoter,Randomer
