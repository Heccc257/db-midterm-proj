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

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word*123456' \
--form 'uid="1"' \
--form 'offer_id="10"'

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

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass_word_bob*54321' \
--form 'uid="2"' \
--form 'offer_id="11"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass5*55555' \
--form 'uid="5"' \
--form 'offer_id="15"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass6*66666' \
--form 'uid="6"' \
--form 'offer_id="17"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass7*77777' \
--form 'uid="7"' \
--form 'offer_id="19"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/accept_offer/pass8*88888' \
--form 'uid="8"' \
--form 'offer_id="13"'

echo '\n'