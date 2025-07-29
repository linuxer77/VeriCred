# VeriCred: Decentralized Digital Credentials

VeriCred is a cutting-edge platform revolutionizing academic and professional credentialing. It leverages blockchain technology (NFTs on Polygon) for secure, transparent, and instantly verifiable digital records of achievements, empowering individuals and streamlining processes for institutions and verifiers.

---

## Key Features & Benefits

- **Forgery-Proof Credentials:** NFTs signed by verified issuers.
- **Instant & Global Verification:** Real-time validation via blockchain.
- **Individual Ownership:** Credentials managed in personal digital wallets.
- **Reduced Admin Burden:** Streamlined issuance for institutions.
- **Enhanced Portability:** Easy sharing of digital achievements.

---

## Technology Stack

- **Blockchain:** Polygon
- **Smart Contracts:** Solidity (ERC-721, IssuerRegistry, RevocationRegistry)
- **Backend:** NestJS
- **Frontend:** React / Next.js
- **Database:** PostgreSQL
- **Storage:** IPFS / Arweave
- **Wallets:** MetaMask, WalletConnect

---

## Getting Started

### Prerequisites

- Node.js & npm/yarn
- Python (for contract tools)
- Docker (optional)
- Polygon RPC endpoint (e.g., Mumbai testnet)
- A crypto wallet (e.g., MetaMask)

### Smart Contracts

1.  Clone the contract repository.
2.  Install dependencies: `npm install`
3.  Compile: `npx hardhat compile`
4.  Deploy to testnet/local network using deployment scripts.

### Backend

1.  Clone the backend repository.
2.  Install dependencies: `npm install`
3.  Configure `.env` file (database, RPC, secrets).
4.  Start server: `npm run start:dev`

### Frontend

1.  Clone the frontend repository.
2.  Install dependencies: `npm install`
3.  Configure `.env` file (backend API URL, RPC).
4.  Start development server: `npm run dev`

---

## Contributing

Contributions are welcome! Please see `CONTRIBUTING.md` for details.

---

## License

This project is licensed under the [MIT License](LICENSE).

---
