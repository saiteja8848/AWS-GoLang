# AWS-GoLang
It is a full stack web project developed  front end using react and backend using go language, aws.


# Running Application
Below is the deployed react-app link,

### Note:   Since I have not get my ssl certificate validated,It will block the access to API. So Please make sure your when you open that link in chrome or firefox or any other browser to allow insecure content for that specific site otherwise it won't see any changes.

### For Google chrome, click site settings, set "allow" for insecure content.
![browser](https://user-images.githubusercontent.com/47274869/99748503-a601d900-2b02-11eb-87ba-ddcd46ce902d.jpg)

 Please checkout App Link: https://react-chart-app.herokuapp.com/ , **Initially, It will display a empty line graph, where you don't see any values. For _every minute_ you can see the changes in the graph since values are been fetched and plotting from the api.**
 
 
 output-1 ![output-1](https://user-images.githubusercontent.com/47274869/99748237-28d66400-2b02-11eb-8d7a-881f29755cee.jpg)
 
 Please observe, so that you can see that values have been mapping on to the graph.
 
 output-2: ![output-2](https://user-images.githubusercontent.com/47274869/99747765-30493d80-2b01-11eb-9d71-c5f66b6eac24.jpg)
 
 get api format
 ![response](https://user-images.githubusercontent.com/47274869/99750445-6ccb6800-2b06-11eb-9634-befc6d0953f8.jpg)
 
 In the above response, I am **using _score as mapping value on to the graph_**.  
 
 http://golangcruddynamodb-env.eba-jcmcekqk.ap-south-1.elasticbeanstalk.com/getCategorys
 

# Steps Followed In Creating React-Chart-App
 prerequisites: node,npm installed with node, create-react-app
 1.create a react app:  ```create-react-app react-chart-app ``` 
 2.Install  rechart for displaying line graph ```npm install --save react-chartjs-2 chart.js``` 
 3.create a component, use rechart for displaying line-chart and fetch, for access api
 4.Run by using ```npm start```


# Steps Followed In Creating Aws-GoLang-Crud-DynamoDB
 1. Install serverless, awscli for creating aws-go project and configure aws in our local system
 2.```npm i -g serverless```  
 3.we will get creditionals (access key, secret key) by creating a user in aws, IAM, use them to configure in awscli.
 4.```serverless config credentials --provider provider --key key --secret secret.```
 5.we can use serverless to create aws-go project template, we can deploy once we done by using 
 6.```serverless deploy --stage "stageName" --region "region name"``` 

Project Structure (I have followed below the blog articles to do the entire project)
![project-structure](https://user-images.githubusercontent.com/47274869/99754642-d0f22a00-2b0e-11eb-8e4b-85114a3d4b01.jpg)

  

# Referernces
1. https://www.softkraft.co/aws-lambda-in-golang/#golang-aws-lambda-function-example
2. https://www.javacodegeeks.com/2018/11/build-restful-api-go-using-aws-lambda.html
3. https://github.com/yosriady/serverless-crud-gohttps://github.com/yosriady/serverless-crud-go
4. https://dev.to/wingkwong/building-serverless-crud-services-in-go-with-dynamodb-part-1-2kec
5. https://www.youtube.com/watch?v=Pa99PT16tmw
6. https://youtu.be/T26V1aSEtJE




 
 
 


