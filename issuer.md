### User Story
As a user Using kubeeye in a production environment, we not only uses OPA rule validation, we also used
1. check nodes ssh conntecion
2. kubernetes certexpire
3. check Component startup configuration consistency，such as kubelet command line parameter --root-dir. 

and so on.
While meeting the above, we also hope to extend it with out of tree.

### Detailed Description
Based on the above points, we extend kubeeye and refactor the code.
#### Feature Description
We have added the following features
##### custom command
Expand kubeeye's command line
##### Embed Rules
>Embedded rules, package the rules into kubeeye for easy use
- OPA rules
- Function rules

Function check rules provide more customized rule checks. For example, by using a shell and calling a third-party interface, you can enclose the function and return the output according to the agreed format, which can be displayed uniformly in the report.

##### Why
###### custom command
On the one hand, kubeeye can be programmed into subcommands of other command-line tools. On the other hand, other command-line tools can also become kubeeye tools.
###### Embed Rules
Checklist are different in different environments and different businesses, But they have something in common, If it is maintained only through an external directory, it will lead to redundancy of the checklist. Therefore, we can package it as a whole and control the start and stop of the business checklist through the configuration file later such as
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: kubeeye-<xxxx>-rules
  namespace: kube-system
data:
  version: "v1"
  regorules: |
    enable: 
      - name: allowPrivilegeEscalationRule
      - name: canImpersonateUserRoleRule
    disable: 
      - name: "*"
  funcrules: |
    enable:
      - name: nodeSSHConnection
    disable:
      - name: xxxxStatus
```
Configuration file feature is still in progress.
##### How
###### custom command
A kubeeyecommand is defined using the builder pattern, You can assemble it with any command, regorule and funcrule，Finally, a cobra command line is returned.
###### Embed Rules
- The OPA rule uses go1.16 embedded, It can package files into code compilation. Whether default rules or additional rules,You must use a variable to package OPA rules.
- Function rule is much simpler, Because it is go code itself, it can be packaged and compiled directly through import.
#### Refactor Description
In order to better add new features, we have adjusted the code structure.
- Added directory funcrules, regorules, register
  
    Funcrules: Storing default function rules

    Regorules: Storing default regorules rules

    Register: rules register
- Use go channel mode, Fan In to connect pipeline in series, The main entrance is audit.Run
- Use fs.FS abstracts local file and embedded file operations
- Simplifies the function of output
- Some channels have been merged
### Anything else you would like to add:
https://github.com/leonharetd/kubeeye is refactor kubeeye code
https://github.com/leonharetd/kubeeye_sample is kubeeye sample
These are some of my practices. Welcome to communicate. Thank you very much🙏.