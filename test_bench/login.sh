curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="123456"' \
--form 'password="pass_word"'

echo '\n'

# wrong phone
curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="123456xxxx"' \
--form 'password="pass_word"'

echo '\n'

# wrong pass
curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="123456"' \
--form 'password="pass_wordxxxx"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="54321"' \
--form 'password="pass_word_bob"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="999999"' \
--form 'password="pass_word_carol"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="55555"' \
--form 'password="pass5"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="66666"' \
--form 'password="pass6"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="77777"' \
--form 'password="pass7"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/login' \
--form 'phone_number="88888"' \
--form 'password="pass8"'

echo '\n'