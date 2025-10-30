# GoVite - Go SSR + Vite + TypeScript

A server-side rendered Go application with Vite for frontend asset bundling and TypeScript support.

## Project Structure

```
govite/
├── assets/          # Static assets (images, fonts, etc.)
├── dist/            # Production build output (generated)
├── handlers/        # Go HTTP handlers
│   └── home.go
├── ui/              # Frontend code
│   ├── components/  # UI components
│   │   └── header/
│   │       ├── header.gohtml  # Go template
│   │       └── header.css     # Component styles
│   ├── index.gohtml    # Main page template
│   ├── main.ts         # TypeScript entry point
│   ├── package.json    # Node dependencies
│   ├── tsconfig.json   # TypeScript configuration
│   └── vite.config.ts  # Vite configuration
├── main.go          # Go server entry point
└── go.mod           # Go module definition
```

## Setup

### Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- npm or yarn

### Installation

1. Install Node dependencies:
```bash
cd ui
npm install
cd ..
```

2. Initialize Go module (already done):
```bash
go mod tidy
```

## Development

For development, you need to run both the Vite dev server and the Go server:

### Terminal 1: Start Vite dev server
```bash
cd ui
npm run dev
```

This will start Vite on `http://localhost:5173` with hot module replacement (HMR).

### Terminal 2: Start Go server in dev mode
```bash
go run main.go -dev
```

Or:
```bash
DEV_MODE=true go run main.go
```

The Go server will run on `http://localhost:8080` and serve templates that load assets from the Vite dev server.

Visit `http://localhost:8080` in your browser.

## Production

### Build

1. Build frontend assets:
```bash
cd ui
npm run build
cd ..
```

This creates optimized bundles in the `dist/` directory with a `manifest.json` file.

2. Build Go binary:
```bash
go build -o govite
```

### Run

```bash
./govite
```

Or:
```bash
go run main.go
```

The server will run on `http://localhost:8080` (default) and serve pre-built assets from the `dist/` directory.

### Custom Port

```bash
./govite -port=3000
```

## How It Works

### Development Mode

- Vite dev server runs on port 5173
- Go templates include references to `http://localhost:5173/@vite/client` and `http://localhost:5173/main.ts`
- Vite provides hot module replacement (HMR) for instant updates
- Go server runs on port 8080 and serves SSR templates

### Production Mode

- Vite builds optimized bundles with content hashes
- A `manifest.json` file maps source files to their hashed output files
- Go reads the manifest and injects correct asset paths into templates
- Go server serves static files from the `dist/` directory

### Adding Components

1. Create a new component directory in `ui/components/`:
```
ui/components/mycomponent/
├── mycomponent.gohtml
└── mycomponent.css
```

2. Define the template in `mycomponent.gohtml`:
```html
{{define "mycomponent"}}
<div class="mycomponent">
  <!-- component markup -->
</div>
{{end}}
```

3. Import the CSS in [ui/main.ts](ui/main.ts):
```typescript
import './components/mycomponent/mycomponent.css'
```

4. Use the component in any template:
```html
{{template "mycomponent" .}}
```

## TypeScript

TypeScript files are located in the `ui/` directory. The entry point is [ui/main.ts](ui/main.ts).

Path aliases are configured:
- `@/` maps to `ui/`
- `@components/` maps to `ui/components/`

Example:
```typescript
import '@components/header/header.css'
```

## Environment Variables

- `DEV_MODE=true` - Run in development mode (alternative to `-dev` flag)
- `PORT=3000` - Alternative to `-port` flag

## Notes

- All frontend code should live in the `ui/` directory
- Component templates use `.gohtml` extension for Go templates
- Static assets (images, fonts) go in the `assets/` directory
- The Go server handles all routing and SSR
- Vite only handles asset bundling and development HMR
