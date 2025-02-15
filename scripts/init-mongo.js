db = db.getSiblingDB("insider_db");

db.messages.insertMany([
  {
    phone_number: "+905551111111",
    content: "First Message for Insider Assessment Project",
    is_sent: false,
    created_at: new Date(),
    sent_at: null 
  },
  {
    phone_number: "+905552222222",
    content: "Second Message for Insider Assessment Project",
    is_sent: false,
    created_at: new Date(),
    sent_at: null   
  },
  {
    phone_number: "+905553333333",
    content: "Third Message for Insider Assessment Project",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905554444444",
    content: "Message 4",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905555555555",
    content: "Message 5",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905553333333",
    content: "Mesaj 3",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905553333333",
    content: "Mesaj 3",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905556666666",
    content: "Message 6",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905557777777",
    content: "Message 7",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905558888888",
    content: "Message 8",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905559999999",
    content: "Message 9",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905550000000",
    content: "I am sending a message that will exceed the character limit. This message will not be processed. We set the character limit to 160. Messages over this limit will not be processed.",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905551111111",
    content: "Message 11",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905552222222",
    content: "Message 12",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905553333333",
    content: "Message 13",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905554444444",
    content: "Message 14",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905555555555",
    content: "Message 15",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905556666666",
    content: "Message 16",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905557777777",
    content: "Message 17",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905553333333",
    content: "Mesaj 3",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905553333333",
    content: "Mesaj 3",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905558888888",
    content: "Message 18",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905559999999",
    content: "Message 19",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905550000000",
    content: "Message 20",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905551111111",
    content: "Message 21",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905552222222",
    content: "Message 22",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905553333333",
    content: "Message 23",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905554444444",
    content: "Message 24",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905555555555",
    content: "Message 25",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905556666666",
    content: "Message 26",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905557777777",
    content: "Message 27",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905558888888",
    content: "Message 28",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905559999999",
    content: "Message 29",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  },
  {
    phone_number: "+905550000000",
    content: "Message 30",
    is_sent: false,
    created_at: new Date(),
    sent_at: null
  }
]);

print("MongoDB: Initial messages inserted.");
