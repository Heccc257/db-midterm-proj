curl --location --request POST 'http://127.0.0.1:9999/offer/post' \
--form 'reward_amount="200.0"' \
--form 'customer_id="1"' \
--form 'category_name="外卖"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="2023-05-05 17:29:02.295078175"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post' \
--form 'customer_id="1"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="2023-05-05 17:29:02.295078175"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post' \
--form 'reward_amount="300.0"' \
--form 'customer_id="2"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="2023-05-05 17:29:02"'

echo '\n'

# 检验与User的关联是否正确
curl --location --request POST 'http://127.0.0.1:9999/offer/post' \
--form 'reward_amount="300.0"' \
--form 'customer_id="1000000"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="2023-05-05 17:29:02"'

echo '\n'

# 没有category
curl --location --request POST 'http://127.0.0.1:9999/offer/post' \
--form 'reward_amount="300.0"' \
--form 'customer_id="2"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="2023-05-05 17:29:02"'

echo '\n'

# 没有location
curl --location --request POST 'http://127.0.0.1:9999/offer/post' \
--form 'reward_amount="1300.0"' \
--form 'customer_id="2"' \
--form 'category_name="跑腿"' \
--form 'time_limit="2023-05-05 17:29:02"'

echo '\n'