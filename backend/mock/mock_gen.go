package mock

//go:generate mockgen -destination=../domain/repository/mock/user_repository.go github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository IfUserRepository
//go:generate mockgen -destination=../domain/repository/mock/article_repository.go github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository IfArticleRepository
//go:generate mockgen -destination=../domain/externals/gcp/mock/firebase_client.go github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/gcp IfFirebaseClient
//go:generate mockgen -destination=../domain/externals/datadog/mock/datadog_client.go github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/datadog IfDatadogClient
//go:generate mockgen -destination=../domain/externals/newrelic/mock/newrelic_client.go github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/externals/newrelic IfNewrelicClient
