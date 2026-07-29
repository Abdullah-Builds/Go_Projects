docker run -d `
  --name mailpit `
  --restart unless-stopped `
  -p 8025:8025 `
  -p 1025:1025 `
  axllent/mailpit