1. Struct Subscription this is domain entity [Layer Violations 1 ]
2. Usecase (bussines logic) purity [Layer Violations 2]
3. Subscription Service knows about http.Client and can't do direct call [Layer Violations 3 ] must be dedicate layer
4. Database execution called within CreateSubscription of Subscription Service   [Layer Violations 4 ]
5. Forming HTTP payload in serice logic [Layer Violations 5]


1. Struct Subscription has Public properties Anemic Domain Model  [Domain Purity 1 ]
2. Сервіс знає що в яксоті бази даних він буде викорстовувати [DOMAIN PURITY 2]
3. Calculate refund must be next to entity. Model is smart takes responsibility for its own behavior.

    // Calculate refund
    daysUsed := time.Since(sub.StartDate).Hours() / 24
    refundAmount := sub.Price * (30 - daysUsed) / 30

1. Subscription Service hard to test we can't mock external services properly [ Testability 1 ] MUST USE layers interface
for proper segregation and ability to mock those layers.
2. Subscription Service hard to test we can't mock database calls [ Testability 2 ] MUST USE layers interface
for proper segregation and ability to mock those layers.
3. Hardcoded time.Now, time.Since

1. When we call billing API no error handling hard to track what went wrong in logs [Error handling]
    resp, _ := s.client.Get("https://api.billing.com/validate/" + req.CustomerID)
2. No Error handling while json Decoding
3. No Error hadnlign while Scaning database in CancelSubscription
4. No Error handling while making refund


1. Subscription Struct has Price property with float64 type which leads
    to money loses eventually due to long numbers after the dot [Money Handling] Turn into cents and int64
2. Wrap Cancel subscription in transaction



???
* what is it?
    var result struct{ Valid bool }
    json.NewDecoder(resp.Body).Decode(&result)

*     if !result.Valid {
          return nil, errors.New("invalid customer") -> is if enough to handle error?
      }

* What about CreateRequest validation?

*     sub := &Subscription{
          ID:         uuid.New().String(),
          CustomerID: req.CustomerID,
          PlanID:     req.PlanID,
          Price:      req.Price,
          Status:     "ACTIVE", ---> add constat or even ENUM with statues
          StartDate:  time.Now(),
      }

*     if sub.Status == "CANCELLED" {
          return errors.New("already cancelled")
      }

            if sub.IsCancelled() {
                return errors.New("already cancelled")
            }


    // Call refund API
    Make background process with retry logic and then somehow notify
     CancelSubscription method to update subscription row in database with new status

    External calls in service layer or event-driven


