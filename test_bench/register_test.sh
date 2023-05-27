echo "---------------- register --------------\n"

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="SuperAlice"' \
--form 'full_name="Alice"' \
--form 'phone_number="123456"' \
--form 'password_hash="pass_word"'

echo '\n'

# 没有nickname
curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name=""' \
--form 'full_name="Alice"' \
--form 'phone_number="123456"' \
--form 'password_hash="pass_word"'
echo '\n'

# 重复的电话号码
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

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="aaa"' \
--form 'full_name="aaa"' \
--form 'phone_number="111"' \
--form 'password_hash="111"'
echo '\n'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="user5"' \
--form 'full_name="user5"' \
--form 'phone_number="55555"' \
--form 'password_hash="pass5"'
echo '\n'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="user6"' \
--form 'full_name="user6"' \
--form 'phone_number="66666"' \
--form 'password_hash="pass6"'
echo '\n'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="user7"' \
--form 'full_name="user7"' \
--form 'phone_number="77777"' \
--form 'password_hash="pass7"'
echo '\n'

curl --location --request POST '127.0.0.1:9999/register' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)' \
--form 'nick_name="user8"' \
--form 'full_name="user8"' \
--form 'phone_number="88888"' \
--form 'password_hash="pass8"'
echo '\n'