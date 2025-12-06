{
  description = "Go development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go compiler and tools
            go
            
            # Go development tools
            gopls          # Go language server
            golangci-lint  # Go linter
            delve          # Go debugger
            go-tools       # Additional Go tools (goimports, etc.)
            
            # Build tools
            gnumake
            gcc
            
            # Version control
            git
            
            # Additional utilities
            jq             # JSON processor
            curl           # HTTP client
            
            # Optional: Database tools if needed
            # postgresql
            # redis
            # sqlite
          ];

          shellHook = ''
            echo "🐹 Go development environment loaded!"
            echo "Go version: $(go version)"
            echo ""
            echo "Available tools:"
            echo "  go        - Go compiler"
            echo "  gopls     - Language server"
            echo "  golangci-lint - Linter"
            echo "  dlv       - Delve debugger"
            echo "  goimports - Import formatter"
            echo ""
            echo "Quick start:"
            echo "  go mod init <module-name>"
            echo "  go run main.go"
            echo "  go build"
            echo "  go test ./..."
            
            # Set up Go environment
            export GOPATH=$PWD/go
            export PATH=$GOPATH/bin:$PATH
            
            # Create go directory if it doesn't exist
            mkdir -p $GOPATH/bin
            mkdir -p $GOPATH/src
            mkdir -p $GOPATH/pkg
          '';

          # Environment variables
          CGO_ENABLED = "1";
          GO111MODULE = "on";
        };

        # Optional: You can also define packages to build
        packages.default = pkgs.buildGoModule {
          pname = "my-go-app";
          version = "0.1.0";
          src = ./.;
          
          # You'll need to set this when you have dependencies
          # vendorSha256 = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=";
          vendorHash = null; # Use this for apps with no dependencies
        };
      });
}
