# Dummy Record Generator
This is an application written as part of the learning Go experience. It utilises the ['faker'](https://github.com/manveru/faker) library, to display a list of 'person' records, which are defined by the user.
## Application Operation
1. Type 'go run main.go' at the terminal prompt.
2. Enter the number of records to be generated, at the following prompt.

**Enter the number of dummy records ('q' or 'Q' to quit):**

⋅⋅⋅Entering 'q' or 'Q' at the prompt will quit the application. Entering anything non-numeric, will display the following.

**Entered number is not numeric. Please retry.**  
**Enter the number of dummy records ('q' or 'Q' to quit):**

3. Entering a number will action the application to produce and display the required number of dummy records. The application will also display the number of records requested, and the number of successfully created records.

**Enter the number of dummy records ('q' or 'Q' to quit): 2**  

**Record 1**  
**---------**  
**Name          Ms. Bridie Runolfsdottir**  
**Email         dee.pfeffer@reynolds.info**  
**Street:       56224 Kovacek Rapid**  
**City:         Leschtown**  
**Postcode:     67846**  
**State:        New Mexico**  
**Country:      Fiji**  
**Home Phone:   480.225.2051 x0771**  
**Mobile Phone: (634)655-3867 x30741**  

**Record 2**  
**---------**  
**Name          Ms. Miss. Kolby Nitzsche**  
**Email         drew@walterbashirian.net**  
**Street:       57324 Alanis Wall**  
**City:         Halvorsonhaven**  
**Postcode:     80225-2051**  
**State:        New Hamphire**  
**Country:      Bermuda**  
**Home Phone:   (375)863-3188**  
**Mobile Phone: 1-852-108-5601**  

**Records requested:                2**  
**Records successfully produced:    2**

