{
  description = "A set of simple tools to use with sveltekit";

  inputs.nixpkgs.url = "nixpkgs/nixos-24.05";

  outputs = { self, nixpkgs }:
    let
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;

      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.buildGoModule {
            pname = "sveltool";
            inherit version;
            src = ./.;

            vendorHash = "sha256-Tklz8B/lgvrQV4pBtALDG9F9FGZBIEY/9Quc7mAKhMs=";
          };
        });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            packages = with pkgs; [
              go
              cobra-cli
            ];
          };
        });

      overlays.default = final: _prev: {
        jvdm.sveltool = self.packages.${final.system}.default;
      };
    };
}

