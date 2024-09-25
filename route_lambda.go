const AWS = require('aws-sdk');
const lambda = new AWS.Lambda();

exports.handler = async (event) => {
  const randomNumber = Math.random();
  let targetLambda;
  
  if (randomNumber < 0.2) {
    targetLambda = 'airports_v2_function'; // 20% traffic to v2
  } else {
    targetLambda = 'airports_v1_function'; // 80% traffic to v1
  }
  
  const params = {
    FunctionName: targetLambda,
    Payload: JSON.stringify(event),
  };
  
  const response = await lambda.invoke(params).promise();
  return JSON.parse(response.Payload);
};
