# Bellz Culture — E-commerce Platform

A secure, reliable and customer-focused e-commerce platform for Bellz Culture, a fashion and handmade products business specialising in personally crafted bags, footwear, clothing and accessories.

The platform is being developed in Go (Golang), with a focus on handmade product sales, personalised customer orders, inventory management, secure online payments and order fulfilment.

---

## 1. Project Overview

Bellz Culture is being developed to provide customers with a convenient way to discover, customise and purchase products online.

The business focuses primarily on personally handmade products while also offering ready-made products.

The platform will support two main purchasing experiences:

- Standard purchases of available ready-made products.
- Handmade and personalised orders that may require customer specifications, business review, quotation and production.

The system will provide the business owner with administrative tools for managing products, stock, customer orders, production and delivery.

## 2. Project Objectives

The main objectives are to:

1. Establish a professional online presence for Bellz Culture.
2. Showcase handmade and ready-made products.
3. Allow customers to purchase products online.
4. Support product variations such as colours and sizes.
5. Allow customers to request personalised designs and modifications.
6. Provide a quotation and approval workflow for custom orders.
7. Support full payment and deposit-plus-balance arrangements.
8. Integrate secure online payments through Paystack.
9. Provide customers with order tracking.
10. Support order notifications through WhatsApp and email.
11. Provide an administrative dashboard for business operations.
12. Protect customer information, business records and payment processes.

## 3. Business Scope

### Product Categories

The platform will support products including:

- Handmade bags.
- Footwear.
- Clothing.
- Fashion accessories.
- Other products approved by the business owner.

### Product Fulfilment Types

**Ready-made products**

Products that have already been produced and are available for purchase. Inventory quantities and availability must be managed.

**Made-to-order products**

Products manufactured after a customer places an order. The business will define production requirements and estimated preparation times.

**Personalised products**

Customers can submit requests for new designs or modifications to existing designs. These requests require business review and an approved quotation before becoming confirmed orders.

## 4. Core Functional Requirements

### Customer Storefront

- Responsive homepage and navigation.
- Product categories and product listings.
- Product details, descriptions and photographs.
- Product search and filtering.
- Product variants, including applicable colours and sizes.
- Product availability information.
- Shopping cart.
- Customer checkout.

### Personalised Order Management

- Custom design request forms.
- Customer descriptions and specifications.
- Reference image uploads.
- Business review of submitted requests.
- Quotation preparation.
- Customer quotation approval.
- Agreed production time and payment terms.
- Custom order conversion into a confirmed order.

### Payment Management

Paystack will be the planned online payment provider.

The system will support:

- Full upfront payment.
- Deposit followed by balance payment.
- Payment verification.
- Payment status records.
- Outstanding balance tracking.
- Payment confirmation and reconciliation.

The business owner will determine the payment arrangement for each approved personalised quotation.

### Order Management

- Unique order references.
- Order records and order history.
- Order status management.
- Production progress.
- Order cancellation and exception handling.
- Customer order tracking.

### Customer Order Tracking

Customers will be able to check their orders on the website using secure order-tracking credentials.

Tracking information will distinguish between payment, production, dispatch and delivery stages.

### Notifications

The platform will support order-related notifications through:

- WhatsApp.
- Email.

Notifications should communicate relevant order changes, payment confirmations and delivery updates.

### Administrative Dashboard

The business owner will be able to manage:

- Products and categories.
- Product photographs and descriptions.
- Prices and product variants.
- Inventory and availability.
- Custom design requests.
- Quotations and payment terms.
- Customer orders.
- Production progress.
- Payment records.
- Delivery status.

Administrative access must be protected.

## 5. Proposed Technology Stack

| Component | Technology |
|---|---|
| Backend | Go (Golang) |
| Frontend | HTML, CSS and JavaScript |
| Templates | Go `html/template` |
| Database | PostgreSQL (planned) |
| Payment provider | Paystack |
| Customer notifications | WhatsApp and email integrations |
| Version control | Git and GitHub |
| Hosting | To be selected after deployment assessment |

The initial architecture will favour a well-structured Go application rather than unnecessary distributed services.

Technology choices may be reviewed where there is a clear technical or business reason.

## 6. Security and Reliability Requirements

Security and reliability are fundamental project requirements.

The implementation must account for:

- Secure administrative authentication.
- Password hashing.
- Role-based access control where required.
- Server-side input validation.
- Protection against SQL injection.
- Cross-site scripting (XSS) protection.
- Cross-site request forgery (CSRF) protection.
- Secure session management.
- Rate limiting and brute-force protection.
- Secure file upload validation.
- HTTPS in production.
- Environment-based secret management.
- Secure payment verification.
- Paystack webhook signature verification.
- Duplicate payment and webhook protection.
- Database transaction integrity.
- Safe inventory updates under concurrent requests.
- Error handling and application logging.
- Database backups and recovery procedures.
- Protection against unauthorised access to customer orders.

Sensitive information, passwords, payment secrets and authentication tokens must never be committed to the Git repository.

## 7. Important Business Rules

The application must enforce clear business rules.

1. A custom design request is not automatically a confirmed order.
2. Custom quotations must be reviewed and approved by the business.
3. Customers must accept the quotation and payment terms before proceeding.
4. Full payment or deposit-plus-balance terms must be recorded for each approved custom order.
5. Outstanding balances must be settled according to the agreed business terms before dispatch.
6. Ready-made inventory must not be oversold.
7. Made-to-order products must have realistic production estimates.
8. Product prices and order totals must be calculated and validated on the server.
9. Payment success must be established through trusted server-side verification.
10. Order status changes must follow valid business workflows.
11. Customers must only access their own order information.
12. Delivery must not be marked completed without an authorised status update.

## 8. Edge Cases and Failure Handling

The system must be tested against unexpected situations, including:

- Duplicate checkout submissions.
- Concurrent purchases of the last available item.
- Expired or abandoned payment sessions.
- Successful payment with a failed browser redirect.
- Delayed or duplicated payment webhooks.
- Payment confirmation when the database is temporarily unavailable.
- Incorrect payment amounts.
- Partial payment and outstanding balances.
- Invalid or oversized file uploads.
- Missing customer information.
- Invalid product variants.
- Product price changes during checkout.
- Interrupted network requests.
- Email or WhatsApp service failures.
- Database connection failures.
- Invalid administrative operations.
- Unauthorised order-tracking attempts.

Failures must be handled safely without corrupting orders, losing confirmed payments or exposing customer information.

## 9. Development Roadmap

The project will be developed incrementally.

| Phase | Main deliverables |
|---|---|
| 1. Requirements | Complete SRS and business rules |
| 2. Architecture | Project structure and database design |
| 3. Foundation | Configuration, database, migrations, testing and security |
| 4. Storefront | Responsive customer-facing website |
| 5. Product management | Products, variants and inventory |
| 6. Custom orders | Design requests, quotations and approval |
| 7. Shopping and checkout | Cart, customer details and order creation |
| 8. Payments | Paystack integration and verification |
| 9. Administration | Dashboard and operational management |
| 10. Tracking | Customer tracking and notifications |
| 11. Quality assurance | Functional, security, failure and recovery testing |
| 12. Deployment | Production hosting and launch preparation |

Each phase must be reviewed and tested before moving to the next major phase.

## 10. Development Principles

The following principles will guide implementation:

- Security by design.
- Simplicity without sacrificing reliability.
- Clear separation of responsibilities.
- Maintainable and readable Go code.
- Consistent error handling.
- Automated testing.
- Database integrity.
- Accurate business records.
- No hardcoded production secrets.
- No unverified payment confirmations.
- No unnecessary technical complexity.

## 11. Current Project Status

**Current stage:** Requirements and architecture planning.

The existing Go application provides an initial product catalogue and WhatsApp-based enquiry experience.

The Go module is configured, the application builds successfully, and the existing static product images are available.

However, the following capabilities are not yet implemented or verified:

- PostgreSQL integration.
- Automated application tests.
- Full product and inventory management.
- Shopping cart and checkout.
- Custom-order quotation management.
- Customer order persistence.
- Paystack integration.
- Secure order tracking.
- Administrative dashboard.
- Automated customer notifications.

These features will be implemented and tested progressively.

## 12. Local Development

### Prerequisites

- Go 1.22.2 or a compatible newer version.
- Git.

### Clone the repository

```bash
git clone https://github.com/antenyijoseph20-glitch/bellz_culture.git
cd bellz_culture
```

### Run the application

```bash
go run .
```

By default, the application uses port `8080`.

Open:

```text
http://localhost:8080
```

To use a different port:

```bash
PORT=8081 go run .
```

### Run tests

```bash
go test ./...
```

### Build the application

```bash
go build -o bellz_culture .
```

## 13. Project Repository

GitHub:

https://github.com/antenyijoseph20-glitch/bellz_culture

## 14. Project Commitment

Bellz Culture will be developed as a dependable business platform, not merely a demonstration website.

Every major feature must be reviewed for functionality, security, usability, maintainability and failure handling before being considered ready for production.

**The goal is to deliver a professional e-commerce platform that supports Bellz Culture's handmade business, protects its customers and supports sustainable business operations.**