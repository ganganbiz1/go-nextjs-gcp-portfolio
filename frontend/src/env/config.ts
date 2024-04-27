interface EnvConfig {
  apiBaseUrl: string;
}

const config: EnvConfig = {
  apiBaseUrl: process.env.NEXT_PUBLIC_API_HOST || "http://localhost:9001",
};

export default config;
