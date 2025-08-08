{
  description = "Generate deterministic ed25519 SSH keypairs from mnemonic phrases using BIP39/BIP32 standards";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs =
    { self, nixpkgs }:
    let

      # to work with older version of flakes
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";

      # Generate a user-friendly version number.
      version = builtins.substring 0 8 lastModifiedDate;

      # System types to support.
      supportedSystems = [
        "x86_64-linux"
        "x86_64-darwin"
        "aarch64-linux"
        "aarch64-darwin"
      ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in
    {
      packages = forAllSystems (system: {
        default = nixpkgsFor.${system}.buildGoModule {
          pname = "mnemonic-ssh";
          inherit version;
          src = ./.;
          vendorHash = "sha256-rRLyZ+7aqNd8nNdpa2RTauDhbDwBcHbDrSLlsGaTMGU=";
        };
      });
    };
}
