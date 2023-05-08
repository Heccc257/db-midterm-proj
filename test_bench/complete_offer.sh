curl --location --request PUT 'http://127.0.0.1:9999/complete_offer?uid=1&offer_id=1'

echo '\n'

curl --location --request PUT 'http://127.0.0.1:9999/complete_offer?uid=2&offer_id=2'

echo '\n'

# 不存在的uid
curl --location --request PUT 'http://127.0.0.1:9999/complete_offer?uid=100&offer_id=3'

echo '\n'

# 不存在的offer_id
curl --location --request PUT 'http://127.0.0.1:9999/complete_offer?uid=1&offer_id=100'

echo '\n'