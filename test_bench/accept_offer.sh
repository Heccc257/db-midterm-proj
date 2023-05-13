curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word_bob*54321?uid=2&offer_id=1'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word_bob*54321?uid=2&offer_id=2'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word*123456?uid=1&offer_id=3'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word_bob*54321?uid=2&offer_id=4'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word*123456?uid=1&offer_id=5'

echo '\n'

# 不存在的uid
curl --location --request POST 'http://127.0.0.1:9999/accept_offer/access?uid=100&offer_id=3'

echo '\n'

# 不存在的offer_id
curl --location --request POST 'http://127.0.0.1:9999/accept_offer/access?uid=1&offer_id=100'

echo '\n'