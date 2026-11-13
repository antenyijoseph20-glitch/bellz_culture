# Bellz Culture — E-commerce Platform

A secure, reliable and customer-focused e-commerce platform for Bellz Culture, a fashion and handmade products business specialising in personally crafted bags, footwear, clothing and accessories.

The platform is being developed with Go, server-rendered HTML, CSS, JavaScript and PostgreSQL. It will support ready-made products, made-to-order products, personalised customer orders, inventory management, secure payments and order fulfilment.

---

## 1. Project Vision

Bellz Culture aims to provide a professional online shopping experience while supporting the business's handmade production and personalised-design services.

The platform will support three product fulfilment types:

- **Ready-made products:** Items already produced and available for purchase.
- **Made-to-order products:** Items produced after a customer places an order.
- **Personalised products:** Products customers want designed or modified to their specifications. These require owner review, quotation and customer approval before becoming confirmed orders.

The application will initially operate as a single-business-owner platform. Its architecture should remain maintainable and allow future expansion without introducing unnecessary complexity.

## 2. Project Objectives

1. Establish a professional online storefront.
2. Showcase products with clear photographs, descriptions, prices and availability.
3. Support product categories, colours, sizes and other applicable variants.
4. Allow customers to search, browse and purchase products.
5. Support custom-design requests and reference-image uploads.
6. Provide an owner-managed quotation and approval workflow.
7. Support full payment and deposit-plus-balance arrangements.
8. Integrate Paystack for online payments.
9. Manage orders, production, inventory and delivery.
10. Provide secure customer order tracking.
11. Support relevant WhatsApp and email notifications.
12. Provide a protected administrative dashboard.
13. Protect customer information, business records and payment transactions.

## 3. Agreed Architecture

### 3.1 Architectural approach

The application will use a **modular monolithic architecture**.

This means the application will be one deployable Go application, organised into modules with clearly defined responsibilities.

We will not introduce microservices, Kubernetes or unnecessary distributed infrastructure for the initial business requirements.

### 3.2 Application architecture

| Layer | Responsibility | Technology |
|---|---|---|
| Presentation | Customer pages, forms and browser interactions | HTML, CSS, JavaScript |
| Template rendering | Render dynamic server-side pages | Go `html/template` |
| HTTP layer | Routing, handlers and middleware | Go `net/http` |
| Application layer | Coordinate business operations | Go services |
| Domain layer | Business rules and entities | Go packages |
| Data access | Database queries and persistence | PostgreSQL |
| File storage | Product photographs and customer reference uploads | External filesystem or object storage, as appropriate |
| Payment integration | Payment initiation and verification | Paystack |
| Notifications | Customer communications | WhatsApp and email providers |

### 3.3 Frontend strategy

The initial frontend will use server-rendered HTML templates.

- Go `html/template` will render dynamic pages.
- CSS will control layout, responsiveness and presentation.
- JavaScript will provide browser-side interactions where necessary.
- Templates and CSS will be embedded into the application executable.
- Product photographs and customer-uploaded files will remain outside the executable.

This approach keeps the frontend simple, maintainable and suitable for the business's initial requirements.

### 3.4 Backend strategy

Go will handle:

- HTTP requests and responses.
- Authentication and authorisation.
- Business rules and validation.
- Product and inventory operations.
- Custom-order workflows.
- Cart and checkout processing.
- Order creation and status transitions.
- Payment verification.
- Database transactions.
- Notifications and administrative operations.

Important business decisions must be enforced by the server, not trusted to browser-side JavaScript.

### 3.5 Planned project structure

The current structure is the foundation. Additional packages will be introduced only when their functionality is required.

```text
bellz_culture/
├── cmd/
│   └── web/
│       └── main.go
│
├── internal/
│   ├── app/
│   ├── config/
│   ├── assets/
│   ├── web/
│   │   └── handlers/
│   │
│   ├── platform/
│   │   ├── database/
│   │   ├── security/
│   │   ├── storage/
│   │   └── logging/
│   │
│   ├── catalog/
│   ├── inventory/
│   ├── customers/
│   ├── cart/
│   ├── checkout/
│   ├── customorders/
│   ├── orders/
│   ├── payments/
│   ├── delivery/
│   ├── admin/
│   └── notifications/
│
├── migrations/
├── static/
│   └── images/
├── tests/
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```

**Important:** This is the target structure, not a claim that all these directories and modules already exist. We will create and test them incrementally. Small related modules may be combined where that is simpler and clearer.

## 4. Business Requirements

### 4.1 Product catalogue

The platform will support:

- Handmade bags.
- Footwear.
- Clothing.
- Fashion accessories.
- Other products approved by the business owner.

Products may have:

- Name and description.
- Category.
- Product photographs.
- Price.
- Product variants.
- Availability.
- Fulfilment type.
- Stock quantity where applicable.
- Production estimate where applicable.

### 4.2 Custom-order workflow

A personalised request must follow this process:

1. Customer submits a design request.
2. Customer supplies descriptions, specifications and optional reference images.
3. Business owner reviews the request.
4. Owner prepares a quotation.
5. Customer reviews and accepts or declines the quotation.
6. Owner records the agreed production time and payment arrangement.
7. An approved request becomes a confirmed order.
8. Payment and production proceed according to the agreed terms.

A submitted custom request must never automatically become a confirmed order.

### 4.3 Payment requirements

Paystack is the planned online payment provider.

The platform must support:

- Full upfront payment.
- Deposit followed by balance payment.
- Secure payment initiation.
- Server-side payment verification.
- Payment references and transaction records.
- Outstanding balance tracking.
- Duplicate payment protection.
- Payment reconciliation.

For each approved personalised order, the business owner will choose the payment arrangement.

### 4.4 Order and delivery management

The system will support:

- Unique order references.
- Customer order history.
- Order items and agreed prices.
- Payment status.
- Production status.
- Dispatch and delivery status.
- Order cancellation and exception handling.
- Secure customer tracking.

### 4.5 Administration

The business owner will need protected tools for managing:

- Products, categories and variants.
- Product photographs and prices.
- Inventory and stock movements.
- Custom requests and quotations.
- Customer orders.
- Production progress.
- Payments and outstanding balances.
- Delivery and customer notifications.

## 5. Technology Stack

| Component | Technology |
|---|---|
| Programming language | Go 1.27.4 |
| Frontend | HTML, CSS and JavaScript |
| Templates | Go `html/template` |
| HTTP server | Go `net/http` |
| Database | PostgreSQL |
| Database migrations | To be selected during database implementation |
| Payment provider | Paystack |
| File storage | External filesystem initially; object storage to be assessed |
| Version control | Git and GitHub |
| Testing | Go testing package and additional tools where justified |
| Deployment | To be selected after production requirements are assessed |

Technology changes must have a clear technical or business justification.

---

## 6. Development Roadmap and Acceptance Criteria

Development will follow the sequence below. Each phase must be reviewed and tested before we proceed.

### Phase 1 — Application foundation

**Status: Completed**

Deliverables:

- Go module and application entry point.
- Configuration loading and validation.
- Application factory.
- HTTP router.
- Homepage handler.
- Embedded HTML templates and CSS.
- External product-image serving.
- Graceful server shutdown.
- HTTP server timeouts.
- Initial security headers.
- Automated tests.

Acceptance criteria:

- Application builds successfully.
- Homepage and assets load correctly.
- Invalid image paths and unsupported image types are rejected.
- Security headers are present on application responses.
- Configuration errors are handled safely.
- Tests, race tests and static analysis pass.

### Phase 2 — Customer-facing frontend

**Status: Initial storefront exists; further development planned**

Deliverables:

- Responsive navigation and page layout.
- Homepage.
- Product catalogue.
- Product detail pages.
- Category browsing.
- Search and filtering.
- Product variant selection.
- Availability presentation.
- Custom-order enquiry interface.
- Contact and business information pages.
- Accessible forms and navigation.

Acceptance criteria:

- Pages display correctly on mobile, tablet and desktop.
- Images maintain suitable proportions.
- Forms provide understandable validation feedback.
- Navigation and links work correctly.
- Pages remain usable with JavaScript unavailable where practical.
- Frontend changes do not expose sensitive business information.

### Phase 3 — Database foundation

**Status: Planned**

Deliverables:

- PostgreSQL connection and configuration.
- Database migration strategy.
- Initial schema and constraints.
- Data-access layer.
- Transaction handling.
- Database health checks.

Initial data areas:

- Admin users.
- Customers.
- Categories.
- Products.
- Product variants.
- Product images.
- Inventory and stock movements.
- Custom requests.
- Quotations.
- Orders and order items.
- Payments.
- Delivery records.

Acceptance criteria:

- Migrations can be applied to a clean database.
- Constraints protect important relationships and business rules.
- Database credentials are not hardcoded.
- Database errors are handled without exposing sensitive details.
- Transactions preserve data integrity.

### Phase 4 — Backend catalogue and inventory

**Status: Planned**

Deliverables:

- Product management services.
- Category management.
- Product variants.
- Product availability.
- Inventory tracking.
- Stock movement records.
- Product image management.

Acceptance criteria:

- Invalid products and variants cannot be saved.
- Prices and stock values are validated on the server.
- Stock updates are safe under concurrent requests.
- Ready-made products cannot be oversold.
- Inventory changes are traceable.

### Phase 5 — Customers and custom orders

**Status: Planned**

Deliverables:

- Customer information management.
- Custom-design request submission.
- Secure reference-image uploads.
- Owner review.
- Quotation preparation.
- Customer quotation approval.
- Agreed production estimates.
- Full-payment or deposit-payment terms.

Acceptance criteria:

- A request cannot become an order without the required approval.
- Uploads are validated by content, size and permitted file type.
- Customers cannot access another customer's private request.
- Quotations preserve the agreed price and terms.
- Changes to a quotation are recorded appropriately.

### Phase 6 — Cart and checkout

**Status: Planned**

Deliverables:

- Shopping cart.
- Product and variant selection.
- Customer checkout details.
- Server-side price calculation.
- Order creation.
- Stock reservation or deduction strategy.
- Duplicate submission protection.

Acceptance criteria:

- Browser-submitted prices are never trusted.
- Invalid or unavailable variants cannot be purchased.
- Concurrent purchases cannot oversell stock.
- Duplicate checkout requests do not create duplicate orders.
- Order totals are calculated accurately on the server.

### Phase 7 — Payment integration

**Status: Planned**

Deliverables:

- Paystack account and integration configuration.
- Payment initialisation.
- Payment verification.
- Webhook processing.
- Webhook signature verification.
- Duplicate-event protection.
- Full payment and deposit-plus-balance processing.
- Outstanding balance tracking.
- Payment reconciliation.

Acceptance criteria:

- A browser redirect alone cannot confirm payment.
- The server verifies payment with the payment provider.
- Payment amount, currency and reference are checked.
- Invalid webhook signatures are rejected.
- Duplicate and delayed webhook events are handled safely.
- Payment records and order updates are transactionally consistent.

### Phase 8 — Administrative dashboard

**Status: Planned**

Deliverables:

- Secure administrator login.
- Password hashing.
- Session management.
- Authorisation checks.
- Product and category management.
- Inventory management.
- Custom request and quotation management.
- Order and production management.
- Payment and delivery management.

Acceptance criteria:

- Unauthenticated users cannot access protected administration.
- Unauthorised actions are rejected.
- Administrative input is validated.
- Important actions are logged.
- Session security is configured appropriately.

### Phase 9 — Delivery, tracking and notifications

**Status: Planned**

Deliverables:

- Order tracking.
- Production progress updates.
- Dispatch and delivery updates.
- WhatsApp notifications.
- Email notifications.
- Notification retry and failure handling.

Acceptance criteria:

- Tracking access is restricted to the correct customer or authorised administrator.
- Tracking credentials are difficult to guess.
- Customers cannot modify order status.
- Notification failures do not corrupt orders or payment records.
- Delivery completion requires an authorised update.

### Phase 10 — Security hardening

**Status: Initial HTTP security protections implemented; full review planned**

Security work will cover:

- Password hashing and secure authentication.
- Authorisation and access control.
- SQL injection protection.
- XSS prevention.
- CSRF protection.
- Secure cookies and session handling.
- Rate limiting and brute-force protection.
- Request-size limits.
- Secure file upload validation.
- Path traversal protection.
- Secret management.
- HTTPS and production security headers.
- Paystack verification and webhook security.
- Customer data privacy.
- Secure error handling.
- Logging without exposing secrets.
- Dependency and configuration review.

Acceptance criteria:

- Security controls are tested, not merely configured.
- Sensitive data is not exposed in responses or logs.
- Protected routes reject unauthorised access.
- Upload restrictions are enforced.
- Production secrets are stored outside source control.
- HTTPS is enabled before public production use.

### Phase 11 — Quality assurance and failure testing

**Status: Initial automated tests implemented; comprehensive testing planned**

Testing will include:

1. Unit tests for business rules.
2. HTTP handler and routing tests.
3. Database integration tests.
4. Frontend and browser workflow tests.
5. Payment integration tests using the provider's test environment.
6. Security tests.
7. Concurrency and race tests.
8. Failure and recovery tests.
9. Responsive layout tests.
10. Regression tests before releases.

Important failure scenarios include:

- Duplicate checkout submissions.
- Concurrent purchase of the last available item.
- Expired payment sessions.
- Successful payment with a failed browser redirect.
- Delayed or duplicate webhooks.
- Incorrect payment amounts.
- Partial payments and outstanding balances.
- Invalid or oversized uploads.
- Database connection failures.
- Interrupted network requests.
- Notification service failures.
- Unauthorised tracking attempts.
- Invalid administrative operations.

Acceptance criteria:

- Critical business workflows pass end-to-end tests.
- No known critical security issue remains unresolved.
- Payment and order records remain consistent after simulated failures.
- Recovery and backup procedures are documented and tested.
- Regression tests pass before release.

### Phase 12 — Production deployment

**Status: Planned**

Deliverables:

- Hosting provider selection.
- Production environment configuration.
- PostgreSQL production database.
- Production file storage.
- HTTPS and domain configuration.
- Application build and release process.
- Database migration procedure.
- Logging and monitoring.
- Backup and recovery plan.
- Deployment and rollback documentation.

Acceptance criteria:

- No production secrets are stored in Git.
- HTTPS is active.
- Database backups are scheduled and recovery is tested.
- Application health can be monitored.
- Migrations can be applied safely.
- Errors do not expose internal paths, credentials or customer information.
- Production checkout and payment workflows are verified before public launch.

---

## 7. Security and Reliability Rules

The following rules apply throughout development:

1. All important validation must happen on the server.
2. Passwords must be hashed using an appropriate password-hashing algorithm.
3. Customer and administrator access must be properly authorised.
4. Database queries must use safe parameterised operations.
5. State-changing browser requests must receive appropriate CSRF protection.
6. Payment success must be established through trusted server-side verification.
7. Payment webhooks must be authenticated and processed idempotently.
8. Inventory updates must be safe under concurrent requests.
9. Uploaded files must be validated and stored safely.
10. Customer order information must remain private.
11. Application errors must not reveal secrets or internal implementation details.
12. Database transactions must protect related business operations.
13. Production secrets must remain outside the repository.
14. Backup and recovery must be tested, not assumed.

## 8. Important Business Rules

- A custom request is not a confirmed order.
- The owner must review and quote personalised requests.
- The customer must accept the quotation before confirmation.
- Payment terms must be recorded for each approved custom order.
- Outstanding balances must be settled according to the agreed terms before dispatch, unless the business explicitly approves another arrangement.
- Ready-made inventory must not be oversold.
- Made-to-order products must have realistic production estimates.
- Product prices and order totals must be validated on the server.
- Order status changes must follow valid business workflows.
- Customers must only access their own order information.
- Delivery completion must require an authorised update.

## 9. Development Standards

- Work on one milestone at a time.
- Inspect existing code before modifying it.
- Keep modules small and responsibilities clear.
- Avoid unnecessary dependencies.
- Write tests alongside new functionality.
- Test important failure cases.
- Run formatting, tests and static analysis regularly.
- Review changes before committing.
- Do not push to GitHub without approval.
- Update this README when a milestone is completed or a significant architecture decision changes.

### Required verification commands

```bash
~/go/bin/go1.27.4 fmt ./...
~/go/bin/go1.27.4 test ./...
~/go/bin/go1.27.4 test -race ./...
~/go/bin/go1.27.4 vet ./...
~/go/bin/go1.27.4 build -o /tmp/bellz_culture ./cmd/web
git diff --check
```

A milestone is not complete merely because the code compiles. Its agreed acceptance criteria must also be satisfied.

## 10. Current Project Status

**Current milestone: Phase 1 — Application foundation completed.**

Implemented and verified:

- Go application entry point at `cmd/web/main.go`.
- Configuration loading and validation.
- Application factory.
- HTTP router and homepage handler.
- Embedded HTML templates and CSS.
- External product-image serving.
- Image path and file-type restrictions.
- Symlink escape protection.
- HTTP security headers.
- HTTP server timeouts.
- Graceful shutdown.
- Automated tests for application, configuration, assets, routing and security.

The following verification commands have passed:

- `go test ./...`
- `go test -race ./...`
- `go vet ./...`
- `go build -o /tmp/bellz_culture ./cmd/web`
- `git diff --check`

Existing storefront pages and product images are available.

**Next milestone: Phase 2 — Customer-facing frontend review and development.**

PostgreSQL, full product management, customer accounts, cart, checkout, custom-order management, Paystack, administration, tracking and notifications remain planned work. They must not be treated as implemented until developed and verified.

## 11. Local Development

### Prerequisites

- Go 1.27.4.
- Git.
- Product images in `static/images/`.

### Clone the repository

```bash
git clone https://github.com/antenyijoseph20-glitch/bellz_culture.git
cd bellz_culture
```

### Run the application

```bash
~/go/bin/go1.27.4 run ./cmd/web
```

By default, the server listens on port `8080`.

Open:

```text
http://localhost:8080
```

### Configure the server

The application supports:

| Variable | Default | Purpose |
|---|---|---|
| `PORT` | `8080` | HTTP server port |
| `IMAGE_DIR` | `static/images` | External product-image directory |

Example:

```bash
PORT=8081 IMAGE_DIR="$PWD/static/images" \
  ~/go/bin/go1.27.4 run ./cmd/web
```

### Run tests

```bash
~/go/bin/go1.27.4 test ./...
```

### Build the application

```bash
~/go/bin/go1.27.4 build -o /tmp/bellz_culture ./cmd/web
```

---

## 12. Repository

GitHub:

https://github.com/antenyijoseph20-glitch/bellz_culture

## 13. Project Commitment

Bellz Culture is being developed as a dependable business platform, not merely a demonstration website.

Every major feature must be reviewed for functionality, security, usability, maintainability, data integrity and failure handling before being considered ready for production.

The objective is to deliver a professional e-commerce platform that supports Bellz Culture's handmade business, protects its customers and supports sustainable business operations.