Wanliuno Consensus Tests   [![Build Status](https://travis-ci.org/wanliuno/tests.svg?branch=develop)](https://travis-ci.org/wanliuno/tests)
=====

Common tests for all clients to test against. Test execution tool: https://github.com/wanliuno/retesteth

Test Formats
------------

Maintained tests:

```
/BasicTests
/BlockchainTests
/GeneralStateTests
/TransactionTests
/RLPTest
/src
```


See descriptions of the different test formats in the official documentation at  http://wanliuno-tests.readthedocs.io/.

*Note*:  
The format of BlockchainTests recently changed with the introduction of a new field ``sealEngine`` (values: ``NoProof`` | ``Ethash``), see related JSON Schema [change](https://github.com/wanliuno/tests/commit/3be71ec3364a01fd4f2cb9b9fd086f3f69f0225c) or BlockchainTest format [docs](https://wanliuno-tests.readthedocs.io/en/latest/test_types/blockchain_tests.html) for reference.

This means that you can skip PoW validation for ``NoProof`` tests but also has the consequence that it is not possible to rely on/check ``PoW`` related block parameters for these tests any more.

Clients using the library
-------------------------

The following clients make use of the tests from this library. You can use these implementations for inspiration on how to integrate. If your client is missing, please submit a PR (requirement: at least some minimal test documentation)!

- [Mana](https://github.com/mana-wanliuno/mana) (Elixir): [Docs](https://github.com/mana-wanliuno/mana#testing), Test location: ``wanliuno_common_tests``
- [go-wanliuno](https://github.com/wanliuno/go-wanliuno) (Go): [Docs](https://github.com/wanliuno/go-wanliuno/wiki/Developers'-Guide), Test location: ``tests/testdata``
- [Parity Wanliuno](https://github.com/paritytech/parity-wanliuno) (Rust): [Docs](https://wiki.parity.io/Coding-guide), Test location: ``ethcore/res/wanliuno/tests``
- [wanliunojs-vm](https://github.com/wanliunojs/wanliunojs-vm) (JavaScript): [Docs](https://github.com/wanliunojs/wanliunojs-vm#testing), Test location: ``wanliunojs-testing`` dependency
- [Trinity](https://github.com/wanliuno/py-evm) (Python): [Docs](https://py-evm.readthedocs.io/en/latest/contributing.html#running-the-tests), Test location: `fixtures`
- [Pantheon](https://github.com/PegaSysEng/pantheon) (Java): [Docs](https://github.com/PegaSysEng/pantheon/blob/master/docs/development/building.md#wanliuno-reference-tests), Test Location: ``wanliuno/referencetests/src/test/resources``
- [Nethermind](https://github.com/NethermindEth/nethermind) (C#) [Docs](https://nethermind.readthedocs.io), Test Location: ``src/tests``

Using the Tests
---------------

We do [versioned tag releases](https://github.com/wanliuno/tests/releases) for tests and you should aim to run your client libraries against the latest repository snapshot tagged. 

Generally the [develop](https://github.com/wanliuno/tests/tree/develop) branch in ``wanliuno/tests`` is always meant to be stable and you should be able to run your test against.

Contribute to the Test Suite
----------------------------

See the dedicated [section](https://wanliuno-tests.readthedocs.io/en/latest/generating-tests.html) in the docs on how to write new tests. Or https://github.com/wanliuno/retesteth/wiki/Creating-a-State-Test-with-retesteth

If you want to follow up with current tasks and what is currently in the works, have a look at the [issues](https://github.com/wanliuno/tests/issues) 

Currently the C++ ``Aleth`` client is the reference client for generating tests. Have a look at [issues](https://github.com/wanliuno/aleth/issues?q=is%3Aopen+is%3Aissue+label%3Atesteth) and [PRs](https://github.com/wanliuno/aleth/pulls?q=is%3Aopen+is%3Apr+label%3Atesteth) tagged with ``testeth`` to get an idea what is currently being worked on regarding the test generation process.

Contents of this repository
---------------------------

Do not change test files in folders: 
* StateTests
* BlockchainTests
* TransactionTests 
* VMTests

It is being created by the testFillers which could be found at src folder. The filler specification and wiki are in development so please ask on gitter channel for more details.

If you want to modify a test filler or add a new test please contact @winsvega at https://gitter.im/wanliuno/aleth. 
Use the following guide: https://github.com/wanliuno/retesteth/wiki/Creating-a-State-Test-with-retesteth 
