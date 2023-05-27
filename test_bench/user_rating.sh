curl --location --request POST 'http://127.0.0.1:9999/user_rating/1/pass_word*123456' \
--form 'offer_id="1"' \
--form 'rating="10"' \
--form 'comment="good"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/1/pass_word*123456' \
--form 'offer_id="2"' \
--form 'rating="5"' \
--form 'comment="normal"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/2/pass_word_bob*54321' \
--form 'offer_id="3"' \
--form 'rating="7"' \
--form 'comment="good"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/3/pass_word_carol*999999' \
--form 'offer_id="4"' \
--form 'rating="3"' \
--form 'comment="bad"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/1/pass_word*123456' \
--form 'offer_id="5"' \
--form 'rating="7"' \
--form 'comment="good"'

echo '\n'

# 不正确的offer_id
curl --location --request POST 'http://127.0.0.1:9999/user_rating/2/pass_word_bob*54321' \
--form 'offer_id="100"' \
--form 'rating="5"' \
--form 'comment="normal"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/5/pass5*55555' \
--form 'offer_id="13"' \
--form 'rating="1"' \
--form 'comment="bad"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/6/pass6*66666' \
--form 'offer_id="15"' \
--form 'rating="2"' \
--form 'comment="bad"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/7/pass7*77777' \
--form 'offer_id="17"' \
--form 'rating="4"' \
--form 'comment="normal"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/8/pass8*88888' \
--form 'offer_id="19"' \
--form 'rating="8"' \
--form 'comment="good"'

echo '\n'