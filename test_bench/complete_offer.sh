echo "-------------------- complete offer ----------------------"
curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass_word_bob*54321?uid=2&offer_id=1'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass_word_bob*54321?uid=2&offer_id=2'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass_word_bob*54321?uid=2&offer_id=4'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass_word*123456?uid=1&offer_id=3'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass_word*123456?uid=1&offer_id=5'

echo '\n'

# 不存在的uid
curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/access?uid=100&offer_id=3'

echo '\n'

# 不存在的offer_id
curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass_word*123456?uid=1&offer_id=100'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass8*88888?uid=8&offer_id=13'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass5*55555?uid=5&offer_id=15'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass6*66666?uid=6&offer_id=17'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer/pass7*77777?uid=7&offer_id=19'

echo '\n'