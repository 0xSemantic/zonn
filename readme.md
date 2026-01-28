# Zonn Hub · The Cross-Chain Community Economy Protocol

[![Go Version](https://img.shields.io/badge/Go-1.24.1-blue)](https://golang.org)
[![Cosmos SDK](https://img.shields.io/badge/Cosmos%20SDK-v0.53.4-2c3e50)](https://github.com/cosmos/cosmos-sdk)
[![Ignite CLI](https://img.shields.io/badge/Ignite-v29.7.0-dev-ff6b6b)](https://ignite.com)
[![License](https://img.shields.io/badge/License-Apache%202.0-green)](LICENSE)
[![Phase](https://img.shields.io/badge/Phase-0%20Genesis%20Complete-orange)](https://github.com/0xSemantic/zonn)
[![Contributors](https://img.shields.io/github/contributors/0xSemantic/zonn)](https://github.com/0xSemantic/zonn/graphs/contributors)

**Zonn** is the foundational **social-economic protocol** for Web3, built natively on the Cosmos SDK. It enables any group — DAOs, gaming guilds, creator collectives, investment clubs, local communities — to seamlessly form, fund, govern, and scale **sovereign, tokenized economies** with native cross-chain capabilities via IBC.

Unlike today’s fragmented stack (Discord/Signal + Snapshot + multisig + separate DEX), Zonn delivers a **single, interoperable coordination layer** where social reputation directly drives economic power across blockchains.

### Vision

> “We are not just building a tool; we are planting the flag for a new way to coordinate.”  
> — Levi Chinecherem Chidi (0xSemantic)

The future of organization, work, and value creation is decentralized, community-owned, and cross-chain. Zonn makes launching a digital economy as simple as starting a group chat — democratizing sovereignty, reputation, governance, treasuries, and liquidity for everyone.

### Core Principles

- **Sovereignty** — Every community owns its full stack: tokens, treasuries, governance rules — deployed as self-contained, upgradable CosmWasm contracts.
- **Native Interoperability** — Unified cross-chain identity & reputation via IBC — no bridges, no silos.
- **Reputation-First Coordination** — Contextual, non-transferable, zk-verifiable reputation powers governance, access, and rewards.
- **No-Code Economy Factory** — Launch full-stack economies (DAO, guild, collective…) with templates and one-click deployment.
- **Gas Sharing** — Communities pre-fund “Gas Tanks” to subsidize member transactions → zero-friction onboarding.
- **Privacy & Security by Design** — Selective disclosure via zero-knowledge proofs, audited templates, battle-tested IBC.

### Phase 0 Foundation (Active Development)

- Deterministic wallet-based identity (one primary profile, multi-wallet linking)
- Community contexts (create / join / roles)
- Contextual reputation signals (join, vote, contribute…)
- Governance integration (reputation-weighted voting)
- Treasury management (governance-controlled)
- Canonical event system (append-only, versioned, indexable)

Full technical blueprint & phased roadmap: [docs/Blueprint.md](docs/Blueprint.md)

### Architecture Overview

Zonn Hub is a sovereign **Cosmos SDK application chain** with IBC at its core.

```
Users & Communities
      ↓
Web / Mobile App (React / React Native) → Public API Gateway
      ↓
Zonn Hub (Cosmos SDK chain)
      ↔ IBC ↔ Cosmos Zones (Osmosis, Juno, Stargaze…) → EVM, Solana (future)
```

Core custom modules (under construction):

- `x/identity` — Unified cross-chain profiles
- `x/community` — Sovereign group contexts
- `x/reputation` — Contextual scoring & signals
- `x/treasury` — Governance-controlled funds
- `x/factory` — Economy template engine (CosmWasm)
- `x/gov` — Reputation-weighted proposals & voting

### Quick Start (Local Development)

#### Prerequisites
- Go ≥ 1.24.1
- Ignite CLI ≥ v29 (recommended: latest dev)
- Git

#### Clone & Run
```bash
git clone https://github.com/0xSemantic/zonn.git
cd zonn

# Install dependencies & build
make install

# Start local chain (resets state, funds genesis, starts node)
ignite chain serve --reset-once
```

Endpoints (while running):
- RPC: http://localhost:26657
- API: http://localhost:1317
- Faucet: http://localhost:4500 (for quick test tokens)

See [DEVELOPMENT.md](DEVELOPMENT.md) for detailed setup, testing, debugging, and phase-by-phase progress.

### Building & Running a Full Node

```bash
# Build binary
make build

# Initialize (optional: custom home)
./build/zonnd init zonn-hub --chain-id zonn-1 --home ~/.zonn

# Start node
./build/zonnd start --home ~/.zonn
```

### Contributing

Zonn is open-source and actively seeking early contributors during Phase 0.  
Focus areas right now:

- Implementing & testing core modules (`x/identity`, `x/reputation`, `x/community`…)
- Writing unit & integration tests
- Improving documentation & developer experience
- Helping with Cosmos / Web3 Foundation grant applications

**Contribution Flow**:
1. Fork the repo
2. Create a feature branch (`feature/x-identity-create-profile`)
3. Use Conventional Commits (`feat: add ProfileCreated event`)
4. Open a PR with clear description & test coverage

Please read [CONTRIBUTING.md](CONTRIBUTING.md) and check open issues labeled “good first issue” or “help wanted”.

### Roadmap & Phases

- **Phase 0** (Current) — Chain skeleton, identity, community, reputation basics, treasury, governance integration
- **Phase 1** — Cross-chain reputation signals, IBC hooks, external attestations
- **Phase 2** — Economy Factory, CosmWasm templates, advanced governance (quadratic, conviction)
- **Phase 3** — IBC to major Cosmos zones, external platform anchoring (Discord, GitHub)
- **Phase 4+** — Hybrid DEX (AMM + orderbook), marketplace, mobile app, EVM/Solana expansion

Detailed blueprint: [docs/Blueprint.md](docs/Blueprint.md)

### Connect & Follow Progress

- **Founder & Lead**: Levi Chinecherem Chidi (0xSemantic)  
  - LinkedIn: https://www.linkedin.com/in/0xSemantic  
  - X (Twitter): https://twitter.com/0xSemantic  
  - GitHub: https://github.com/0xSemantic

- **Community** → Coming soon (Discord, Telegram, forum planned for Phase 1+)

- **GitHub Issues & Discussions**: https://github.com/0xSemantic/zonn/issues  
- **Documentation** → In progress (will move to docs.znn.zone)

### License

Apache 2.0 — see [LICENSE](LICENSE)

---

**Zonn is building the missing coordination layer where social capital becomes enforceable economic power — across chains, across communities, across borders.**

Join the journey.

Built with ❤️ by Levi Chinecherem Chidi (0xSemantic) and early contributors.
