import { Wallet } from '../src/wallet';
import { expect } from 'chai';

describe('Wallet', () => {
  let wallet: Wallet;

  beforeEach(() => {
    wallet = new Wallet('test-private-key');
  });

  it('should sign a transaction', () => {
    const tx = { to: '0x...', value: 100 };
    const signedTx = wallet.signTransaction(tx);
    expect(signedTx).to.have.property('signature');
    expect(signedTx.signature).to.match(/^0x[a-fA-F0-9]+$/);
  });

  it('should throw an error for invalid transaction', () => {
    const invalidTx = { to: null, value: -100 };
    expect(() => wallet.signTransaction(invalidTx)).to.throw('Invalid transaction');
  });
});
