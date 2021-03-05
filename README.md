## Graphql Example Application

Domain Entities are:
- Ticket(item) -  an event that must be investigated or a work item that must be addressed.
- Tag(label) - helps you to categorize a ticket, and search Tickets by label later on.
- User(employee) - person that can be assigned to a Ticket.

### Worflow
1. Create Users
2. Create Tags. ex: Urgent, Not-Urgent
3. Create Ticket. Ti
   - Assing User to a Ticket
   - Assign multiple tag

Events - Mutations:
1. UserCreated - CreateUser
2. TagCreated - CreateTag
3. TicketCreated - CreateTicket

Views (queries):
1. Users(pagination, criteria{name, email}) (table view)
2. User(id) (details view)
3. Tags(pagination, criteria{name})
4. Tag(id)
5. Tickets(pagination, criteria{title, assignedUserId, labels})
6. Ticket(id)