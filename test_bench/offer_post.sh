echo "-------------------- post offer ----------------------"
curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="100.0"' \
--form 'customer_id="1"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="100"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/WrongToken' \
--form 'reward_amount="100.0"' \
--form 'customer_id="1"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="120"'

echo '\n'

# 没有reward
curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'customer_id="1"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="50"'

echo '\n'

# 没有reward
curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'customer_id="2"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="30"'

echo '\n'

# 没有time_limit
curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="300.0"' \
--form 'customer_id="1"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="4"' \
--form 'delivery_location_id="1"' \

echo '\n'

# 没有time_limit
curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="300.0"' \
--form 'customer_id="3"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="4"' \
--form 'delivery_location_id="1"' \

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="400.0"' \
--form 'customer_id="2"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="100"'

echo '\n'

# 检验与User的关联是否正确
curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="500.0"' \
--form 'customer_id="1000000"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="60"'

echo '\n'

# 没有category
curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="600.0"' \
--form 'customer_id="2"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="70"'

echo '\n'

# 没有location
curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="700.0"' \
--form 'customer_id="2"' \
--form 'category_name="跑腿"' \
--form 'time_limit="40"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="800.0"' \
--form 'customer_id="3"' \
--form 'pickup_location_id="2"' \
--form 'delivery_location_id="1"' \
--form 'category_name="跑腿"' \
--form 'time_limit="40"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="120.0"' \
--form 'customer_id="1"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="2"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="500"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="220.0"' \
--form 'customer_id="1"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="5"' \
--form 'time_limit="500"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass5*55555' \
--form 'reward_amount="100.0"' \
--form 'customer_id="5"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="2"' \
--form 'delivery_location_id="4"' \
--form 'time_limit="100"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass5*55555' \
--form 'reward_amount="100.0"' \
--form 'customer_id="5"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="2"' \
--form 'delivery_location_id="4"' \
--form 'time_limit="100"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass6*66666' \
--form 'reward_amount="120.0"' \
--form 'customer_id="6"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="2"' \
--form 'delivery_location_id="5"' \
--form 'time_limit="200"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass6*66666' \
--form 'reward_amount="320.0"' \
--form 'customer_id="6"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="5"' \
--form 'delivery_location_id="4"' \
--form 'time_limit="400"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass7*77777' \
--form 'reward_amount="100.0"' \
--form 'customer_id="7"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="3"' \
--form 'delivery_location_id="4"' \
--form 'time_limit="200"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass7*77777' \
--form 'reward_amount="170.0"' \
--form 'customer_id="7"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="3"' \
--form 'delivery_location_id="4"' \
--form 'time_limit="270"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass8*88888' \
--form 'reward_amount="300.0"' \
--form 'customer_id="8"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="4"' \
--form 'delivery_location_id="5"' \
--form 'time_limit="150"'

echo '\n'


curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass8*88888' \
--form 'reward_amount="320.0"' \
--form 'customer_id="8"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="3"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="150"'

echo '\n'

 --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="500.0"' \
--form 'customer_id="2"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="4"' \
--form 'delivery_location_id="1"' \
--form 'time_limit="100"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/pass8*88888' \
--form 'reward_amount="420.0"' \
--form 'customer_id="8"' \
--form 'category_name="拿外卖"' \
--form 'pickup_location_id="3"' \
--form 'delivery_location_id="2"' \
--form 'time_limit="160"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/offer/post/access' \
--form 'reward_amount="300.0"' \
--form 'customer_id="1"' \
--form 'category_name="跑腿"' \
--form 'pickup_location_id="1"' \
--form 'delivery_location_id="5"' \
--form 'time_limit="200"'

echo '\n'