import { act, render, screen } from '@testing-library/react';
import ListCanaries from '../';
import { ProgressiveDeliveryProvider } from '../../../contexts/ProgressiveDelivery';
import {
  defaultContexts,
  ProgressiveDeliveryMock,
  withContext,
} from '../../../utils/test-utils';

describe('ListCanaries', () => {
  let wrap: (el: JSX.Element) => JSX.Element;
  let api: ProgressiveDeliveryMock;

  beforeEach(() => {
    api = new ProgressiveDeliveryMock();
    wrap = withContext([
      ...defaultContexts(),
      [ProgressiveDeliveryProvider, { api }],
    ]);
    api.IsFlaggerAvailableReturns = { clusters: { 'my-cluster': true } };
  });
  it('renders a list of canaries', async () => {
    api.ListCanariesReturns = {
      canaries: [
        {
          name: 'my-canary',
          namespace: 'some-namespace',
          clusterName: 'my-cluster',
          targetReference: {
            kind: 'Deployment',
            name: 'cool-dep',
          },
        },
      ],
    };

    await act(async () => {
      const c = wrap(<ListCanaries />);
      render(c);
    });

    expect(await screen.findByText('my-canary')).toBeTruthy();

    const tbl = document.querySelector('#canaries-list table');
    const rows = tbl?.querySelectorAll('tbody tr');

    expect(rows).toHaveLength(1);
  });
  it('show a link to a promoted image for one container', async () => {
    api.ListCanariesReturns = {
      canaries: [
        {
          name: 'my-canary',
          namespace: 'some-namespace',
          clusterName: 'my-cluster',
          targetReference: {
            kind: 'Deployment',
            name: 'cool-dep',
          },
          targetDeployment: {
            promotedImageVersions: {
              'some-app': 'ghrc.io/myorg/nginx',
            },
          },
        },
      ],
    };

    await act(async () => {
      const c = wrap(<ListCanaries />);
      render(c);
    });

    expect(await screen.findByText('my-canary')).toBeTruthy();

    const promotedCell = findCellInCol('Promoted');
    const text = promotedCell?.textContent;
    expect(text).toContain('ghrc.io/myorg/nginx');
    expect(text).not.toContain('some-app');
  });
  it('shows promoted container images by key/value pairs', async () => {
    api.ListCanariesReturns = {
      canaries: [
        {
          name: 'my-canary',
          namespace: 'some-namespace',
          clusterName: 'my-cluster',
          targetReference: {
            kind: 'Deployment',
            name: 'cool-dep',
          },
          targetDeployment: {
            promotedImageVersions: {
              'some-app': 'ghrc.io/myorg/nginx',
              'other-app': 'ghrc.io/myorg/helloworld',
            },
          },
        },
      ],
    };

    await act(async () => {
      const c = wrap(<ListCanaries />);
      render(c);
    });

    expect(await screen.findByText('my-canary')).toBeTruthy();

    const promotedCell = findCellInCol('Promoted');
    const text = promotedCell?.textContent;
    expect(text).toEqual(
      'some-app: ghrc.io/myorg/nginx other-app: ghrc.io/myorg/helloworld ',
    );
  });
  describe('Canary progress status', () => {
    it('shows correct canary progress when status is below analysis count', async () => {
      api.ListCanariesReturns = {
        canaries: [
          {
            name: 'my-canary',
            namespace: 'some-namespace',
            clusterName: 'my-cluster',
            targetReference: {
              kind: 'Deployment',
              name: 'cool-dep',
            },
            status: {
              phase: 'Progressing',
              iterations: 9,
            },
            analysis: {
              iterations: 10,
            }
          },
        ],
      };

      await act(async () => {
        const c = wrap(<ListCanaries />);
        render(c);
      });

      expect(await screen.findByText('my-canary')).toBeTruthy();

      const statusCell = findCellInCol('Status');
      const text = statusCell?.textContent;
      expect(text).toEqual('9 / 10');
    });
    it('shows correct canary progress when status is equal to analysis count', async () => {
      api.ListCanariesReturns = {
        canaries: [
          {
            name: 'my-canary',
            namespace: 'some-namespace',
            clusterName: 'my-cluster',
            targetReference: {
              kind: 'Deployment',
              name: 'cool-dep',
            },
            status: {
              phase: 'Progressing',
              iterations: 10,
            },
            analysis: {
              iterations: 10,
            }
          },
        ],
      };

      await act(async () => {
        const c = wrap(<ListCanaries />);
        render(c);
      });

      expect(await screen.findByText('my-canary')).toBeTruthy();

      const statusCell = findCellInCol('Status');
      const text = statusCell?.textContent;
      expect(text).toEqual('10 / 10');
    });
    it('shows correct canary progress when status is over analysis count', async () => {
      api.ListCanariesReturns = {
        canaries: [
          {
            name: 'my-canary',
            namespace: 'some-namespace',
            clusterName: 'my-cluster',
            targetReference: {
              kind: 'Deployment',
              name: 'cool-dep',
            },
            status: {
              phase: 'Progressing',
              iterations: 11,
            },
            analysis: {
              iterations: 10,
            }
          },
        ],
      };

      await act(async () => {
        const c = wrap(<ListCanaries />);
        render(c);
      });

      expect(await screen.findByText('my-canary')).toBeTruthy();

      const statusCell = findCellInCol('Status');
      const text = statusCell?.textContent;
      expect(text).toEqual('10 / 10');
    });
  });
});

function findCellInCol(cell: string) {
  const tbl = document.querySelector('#canaries-list table');

  const cols = tbl?.querySelectorAll('thead th');
  const idx = findColByHeading(cols, cell) as number;

  const rows = tbl?.querySelectorAll('tbody tr');

  const promotedCell = rows?.item(0).childNodes.item(idx);

  return promotedCell;
}

// Helper to ensure that tests still pass if columns get re-ordered
function findColByHeading(
  cols: NodeListOf<Element> | undefined,
  heading: string,
): null | number {
  if (!cols) {
    return null;
  }

  let idx = null;
  cols?.forEach((e, i) => {
    if (e.innerHTML.includes(heading)) {
      idx = i;
    }
  });

  return idx;
}