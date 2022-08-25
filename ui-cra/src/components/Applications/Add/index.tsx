import React, { useCallback, useEffect, useMemo, useState } from 'react';
import { ThemeProvider } from '@material-ui/core/styles';
import { localEEMuiTheme } from '../../../muiTheme';
import { PageTemplate } from '../../Layout/PageTemplate';
import { SectionHeader } from '../../Layout/SectionHeader';
import { AddApplicationRequest, useApplicationsCount } from '../utils';
import GitOps from '../../Clusters/Create/Form/Partials/GitOps';
import { Grid } from '@material-ui/core';
import { ContentWrapper } from '../../Layout/ContentWrapper';
import {
  CallbackStateContextProvider,
  clearCallbackState,
  getCallbackState,
  getProviderToken,
} from '@weaveworks/weave-gitops';
import { useHistory } from 'react-router-dom';
import { theme as weaveTheme } from '@weaveworks/weave-gitops';
import { isUnauthenticated, removeToken } from '../../../utils/request';
import useNotifications from '../../../contexts/Notifications';
import { GitProvider } from '@weaveworks/weave-gitops/ui/lib/api/applications/applications.pb';
import { useListConfig } from '../../../hooks/versions';
import { PageRoute } from '@weaveworks/weave-gitops/ui/lib/types';
import AppFields from './form/Partials/AppFields';
import { ClusterAutomation } from '../../../cluster-services/cluster_services.pb';
import useClusters from '../../../contexts/Clusters';
import { useListGitRepos } from '../../../hooks/gitReposSource';
import { Loader } from '../../Loader';

const AddApplication = () => {
  const applicationsCount = useApplicationsCount();
  const [loading, setLoading] = useState<boolean>(false);
  const [showAuthDialog, setShowAuthDialog] = useState(false);
  const history = useHistory();
  const { setNotifications } = useNotifications();
  const { data } = useListConfig();
  const repositoryURL = data?.repositoryURL || '';
  const authRedirectPage = `/applications/create`;
  const { clusters, isLoading } = useClusters();
  const { data: GitRepoResponse } = useListGitRepos();

  const random = useMemo(() => Math.random().toString(36).substring(7), []);

  let initialFormData = {
    url: '',
    provider: '',
    branchName: `add-application-branch-${random}`,
    title: 'Add application',
    commitMessage: 'Add application',
    pullRequestDescription: 'This PR adds a new application',
    clusterAutomations: [
      {
        name: '',
        namespace: '',
        cluster_name: '',
        cluster_namespace: '',
        cluster: '',
        cluster_isControlPlane: false,
        path: '',
        source_name: '',
        source_namespace: '',
        source: '',
      },
    ],
  };

  const callbackState = getCallbackState();

  if (callbackState) {
    initialFormData = {
      ...initialFormData,
      ...callbackState.state.formData,
    };
  }
  const [formData, setFormData] = useState<any>(initialFormData);

  useEffect(() => {
    if (repositoryURL != null) {
      setFormData((prevState: any) => ({
        ...prevState,
        url: repositoryURL,
      }));
    }
  }, [repositoryURL]);

  useEffect(() => {
    clearCallbackState();
  }, []);

  useEffect(() => {
    setFormData((prevState: any) => ({
      ...prevState,
      pullRequestTitle: `Add application ${(formData.clusterAutomations || [])
        .map((a: any) => a.name)
        .join(', ')}`,
    }));
  }, [formData.clusterAutomations]);

  const handleAddApplication = useCallback(() => {
    const clusterAutomations = formData.clusterAutomations.map(
      (kustomization: any): ClusterAutomation => {
        return {
          cluster: {
            name: kustomization.cluster_name,
            namespace: kustomization.cluster_namespace,
          },
          isControlPlane: kustomization.cluster_isControlPlane,
          kustomization: {
            metadata: {
              name: kustomization.name,
              namespace: kustomization.namespace,
            },
            spec: {
              path: kustomization.path,
              sourceRef: {
                name: kustomization.source_name,
                namespace: kustomization.source_namespace,
              },
            },
          },
        };
      },
    );
    const payload = {
      head_branch: formData.branchName,
      title: formData.pullRequestTitle,
      description: formData.pullRequestDescription,
      commit_message: formData.commitMessage,
      clusterAutomations,
    };
    setLoading(true);
    return AddApplicationRequest(
      payload,
      getProviderToken(formData.provider as GitProvider),
    )
      .then(response => {
        history.push('/applications');
        setNotifications([
          {
            message: {
              component: (
                <a
                  style={{ color: weaveTheme.colors.primary }}
                  href={response.webUrl}
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  PR created successfully.
                </a>
              ),
            },
            variant: 'success',
          },
        ]);
      })
      .catch(error => {
        setNotifications([
          { message: { text: error.message }, variant: 'danger' },
        ]);
        if (isUnauthenticated(error.code)) {
          removeToken(formData.provider);
        }
      })
      .finally(() => setLoading(false));
  }, [formData, history, setNotifications]);

  return (
    <ThemeProvider theme={localEEMuiTheme}>
      <PageTemplate documentTitle="WeGo · Add new application">
        <CallbackStateContextProvider
          callbackState={{
            page: authRedirectPage as PageRoute,
            state: {
              formData,
            },
          }}
        >
          <SectionHeader
            className="count-header"
            path={[
              {
                label: 'Applications',
                url: '/applications',
                count: applicationsCount,
              },
              { label: 'Add new application' },
            ]}
          />
          <ContentWrapper>
            <Grid container>
              <Grid item xs={12} sm={10} md={10} lg={8}>
                {!isLoading &&
                  formData.clusterAutomations.map(
                    (kustomization: any, index: number) => {
                      return (
                        <AppFields
                          key={index}
                          formData={formData}
                          setFormData={setFormData}
                          clusters={clusters}
                          GitRepoResponse={GitRepoResponse}
                          index={index}
                        ></AppFields>
                      );
                    },
                  )}
                {isLoading && <Loader></Loader>}
              </Grid>
              <Grid item xs={12} sm={10} md={10} lg={8}>
                <GitOps
                  loading={loading}
                  formData={formData}
                  setFormData={setFormData}
                  onSubmit={handleAddApplication}
                  showAuthDialog={showAuthDialog}
                  setShowAuthDialog={setShowAuthDialog}
                />
              </Grid>
            </Grid>
          </ContentWrapper>
        </CallbackStateContextProvider>
      </PageTemplate>
    </ThemeProvider>
  );
};
export default AddApplication;