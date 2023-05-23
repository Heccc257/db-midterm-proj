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