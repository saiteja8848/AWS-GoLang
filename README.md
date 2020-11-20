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
 1.create a react app, ```create-react-app react-chart-app ``` 



# Steps Followed In Creating Aws-GoLang-Crud-DynamoDB
 1. Install serverless, awscli for creating aws-go project and configure aws in our local system


# Referernces
1. https://www.softkraft.co/aws-lambda-in-golang/#golang-aws-lambda-function-example
2. https://www.javacodegeeks.com/2018/11/build-restful-api-go-using-aws-lambda.html
3. https://github.com/yosriady/serverless-crud-gohttps://github.com/yosriady/serverless-crud-go
4. https://www.youtube.com/watch?v=Pa99PT16tmw
5. https://youtu.be/T26V1aSEtJE




 
 
 


