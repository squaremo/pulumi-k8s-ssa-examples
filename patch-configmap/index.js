"use strict";
const k8s = require("@pulumi/kubernetes");

const provider = new k8s.Provider('ssa', {
    enableServerSideApply: true,
});

const patch = new k8s.core.v1.ConfigMapPatch('foo-patch', {
    metadata: {
        name: 'foo',
        namespace: 'default',
    },
    data: {
        'boo': 'baz',
    },
}, { provider });
