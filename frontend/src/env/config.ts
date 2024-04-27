interface EnvConfig {
  apiBaseUrl: string;
}

const config: EnvConfig = {
  apiBaseUrl: process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:9002",
};

export default config;
