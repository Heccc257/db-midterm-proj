curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nickname="SuperAlice"' \
--form 'full_name="Alice"' \
--form 'phone_number="123456"' \
--form 'password_hash="pass_word"'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nickname=""' \
--form 'full_name="Alice"' \
--form 'phone_number="123456"' \
--form 'password_hash="pass_word"'


curl --location --request GET 'http://127.0.0.1:9999/user_info/1'