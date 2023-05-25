echo "-------------------- accepet offer ----------------------"

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word_bob*54321' \
--form 'uid="2"' \
--form 'offer_id="1"'
echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word_bob*54321' \
--form 'uid="2"' \
--form 'offer_id="2"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word*123456' \
--form 'uid="1"' \
--form 'offer_id="3"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word_bob*54321' \
--form 'uid="2"' \
--form 'offer_id="4"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word*123456' \
--form 'uid="1"' \
--form 'offer_id="5"'

echo '\n'

# 不存在的uid
curl --location --request POST 'http://127.0.0.1:9999/accept_offer/access' \
--form 'uid="100"' \
--form 'offer_id="3"'

echo '\n'

# 不存在的offer_id
curl --location --request POST 'http://127.0.0.1:9999/accept_offer/access?uid=1&offer_id=100' \
--form 'uid="1"' \
--form 'offer_id="100"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word_bob*54321' \
--form 'uid="2"' \
--form 'offer_id="11"'
echo '\n'