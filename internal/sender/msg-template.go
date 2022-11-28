package sender

const MAIL_TEMPLATE = "To: %s\nSubject: Your 'Mikołajki' draw result \nMime-version:1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n<html><head></head><body><div style=\"font-family: Hevletica Verdana; text-align: center\"><p style=\"font-size: 36px\">Hi %s!</p><p style=\"font-size: 24px\">This is your Mikolajkolos!</p><p>Here's who was drawn for you: %s.</p><p>Your drawn friend wrote a short letter to Santa :</p><i>\"%s\"</i></p><p style=\"font-size: 50px\">Ho Ho!</p></div></body></html>"
