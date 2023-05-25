curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="SuperAlice"' \
--form 'full_name="Alice"' \
--form 'phone_number="123456"' \
--form 'password_hash="pass_word"'

echo '\n'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name=""' \
--form 'full_name="Alice"' \
--form 'phone_number="123456"' \
--form 'password_hash="pass_word"'
echo '\n'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="SuperBob"' \
--form 'full_name="bob"' \
--form 'phone_number="123456"' \
--form 'password_hash="pass_word"'
echo '\n'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="SuperBob"' \
--form 'full_name="bob"' \
--form 'phone_number="54321"' \
--form 'password_hash="pass_word_bob"'
echo '\n'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="SuperCarol"' \
--form 'full_name="carol"' \
--form 'phone_number="999999"' \
--form 'password_hash="pass_word_carol"'
echo '\n'


# curl --location --request GET 'http://127.0.0.1:9999/user_info/1'