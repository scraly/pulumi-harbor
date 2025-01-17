# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ReplicationArgs', 'Replication']

@pulumi.input_type
class ReplicationArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 registry_id: pulumi.Input[int],
                 deletion: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dest_namespace: Optional[pulumi.Input[str]] = None,
                 dest_namespace_replace: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationFilterArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override: Optional[pulumi.Input[bool]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 speed: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Replication resource.
        """
        ReplicationArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            registry_id=registry_id,
            deletion=deletion,
            description=description,
            dest_namespace=dest_namespace,
            dest_namespace_replace=dest_namespace_replace,
            enabled=enabled,
            filters=filters,
            name=name,
            override=override,
            schedule=schedule,
            speed=speed,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: pulumi.Input[str],
             registry_id: pulumi.Input[int],
             deletion: Optional[pulumi.Input[bool]] = None,
             description: Optional[pulumi.Input[str]] = None,
             dest_namespace: Optional[pulumi.Input[str]] = None,
             dest_namespace_replace: Optional[pulumi.Input[int]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             filters: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationFilterArgs']]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             override: Optional[pulumi.Input[bool]] = None,
             schedule: Optional[pulumi.Input[str]] = None,
             speed: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'registryId' in kwargs:
            registry_id = kwargs['registryId']
        if 'destNamespace' in kwargs:
            dest_namespace = kwargs['destNamespace']
        if 'destNamespaceReplace' in kwargs:
            dest_namespace_replace = kwargs['destNamespaceReplace']

        _setter("action", action)
        _setter("registry_id", registry_id)
        if deletion is not None:
            _setter("deletion", deletion)
        if description is not None:
            _setter("description", description)
        if dest_namespace is not None:
            _setter("dest_namespace", dest_namespace)
        if dest_namespace_replace is not None:
            _setter("dest_namespace_replace", dest_namespace_replace)
        if enabled is not None:
            _setter("enabled", enabled)
        if filters is not None:
            _setter("filters", filters)
        if name is not None:
            _setter("name", name)
        if override is not None:
            _setter("override", override)
        if schedule is not None:
            _setter("schedule", schedule)
        if speed is not None:
            _setter("speed", speed)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter
    def deletion(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "deletion")

    @deletion.setter
    def deletion(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destNamespace")
    def dest_namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dest_namespace")

    @dest_namespace.setter
    def dest_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dest_namespace", value)

    @property
    @pulumi.getter(name="destNamespaceReplace")
    def dest_namespace_replace(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "dest_namespace_replace")

    @dest_namespace_replace.setter
    def dest_namespace_replace(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dest_namespace_replace", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationFilterArgs']]]]:
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def override(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "override", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def speed(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "speed")

    @speed.setter
    def speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "speed", value)


@pulumi.input_type
class _ReplicationState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 deletion: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dest_namespace: Optional[pulumi.Input[str]] = None,
                 dest_namespace_replace: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationFilterArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override: Optional[pulumi.Input[bool]] = None,
                 registry_id: Optional[pulumi.Input[int]] = None,
                 replication_policy_id: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 speed: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Replication resources.
        """
        _ReplicationState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            deletion=deletion,
            description=description,
            dest_namespace=dest_namespace,
            dest_namespace_replace=dest_namespace_replace,
            enabled=enabled,
            filters=filters,
            name=name,
            override=override,
            registry_id=registry_id,
            replication_policy_id=replication_policy_id,
            schedule=schedule,
            speed=speed,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: Optional[pulumi.Input[str]] = None,
             deletion: Optional[pulumi.Input[bool]] = None,
             description: Optional[pulumi.Input[str]] = None,
             dest_namespace: Optional[pulumi.Input[str]] = None,
             dest_namespace_replace: Optional[pulumi.Input[int]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             filters: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationFilterArgs']]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             override: Optional[pulumi.Input[bool]] = None,
             registry_id: Optional[pulumi.Input[int]] = None,
             replication_policy_id: Optional[pulumi.Input[int]] = None,
             schedule: Optional[pulumi.Input[str]] = None,
             speed: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'destNamespace' in kwargs:
            dest_namespace = kwargs['destNamespace']
        if 'destNamespaceReplace' in kwargs:
            dest_namespace_replace = kwargs['destNamespaceReplace']
        if 'registryId' in kwargs:
            registry_id = kwargs['registryId']
        if 'replicationPolicyId' in kwargs:
            replication_policy_id = kwargs['replicationPolicyId']

        if action is not None:
            _setter("action", action)
        if deletion is not None:
            _setter("deletion", deletion)
        if description is not None:
            _setter("description", description)
        if dest_namespace is not None:
            _setter("dest_namespace", dest_namespace)
        if dest_namespace_replace is not None:
            _setter("dest_namespace_replace", dest_namespace_replace)
        if enabled is not None:
            _setter("enabled", enabled)
        if filters is not None:
            _setter("filters", filters)
        if name is not None:
            _setter("name", name)
        if override is not None:
            _setter("override", override)
        if registry_id is not None:
            _setter("registry_id", registry_id)
        if replication_policy_id is not None:
            _setter("replication_policy_id", replication_policy_id)
        if schedule is not None:
            _setter("schedule", schedule)
        if speed is not None:
            _setter("speed", speed)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def deletion(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "deletion")

    @deletion.setter
    def deletion(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destNamespace")
    def dest_namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dest_namespace")

    @dest_namespace.setter
    def dest_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dest_namespace", value)

    @property
    @pulumi.getter(name="destNamespaceReplace")
    def dest_namespace_replace(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "dest_namespace_replace")

    @dest_namespace_replace.setter
    def dest_namespace_replace(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dest_namespace_replace", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationFilterArgs']]]]:
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ReplicationFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def override(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "override", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter(name="replicationPolicyId")
    def replication_policy_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "replication_policy_id")

    @replication_policy_id.setter
    def replication_policy_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replication_policy_id", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def speed(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "speed")

    @speed.setter
    def speed(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "speed", value)


class Replication(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 deletion: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dest_namespace: Optional[pulumi.Input[str]] = None,
                 dest_namespace_replace: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicationFilterArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override: Optional[pulumi.Input[bool]] = None,
                 registry_id: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 speed: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        Harbor project can be imported using the `replication id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/replication:Replication main /replication/policies/1 <break>```<break><break>`

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        Harbor project can be imported using the `replication id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/replication:Replication main /replication/policies/1 <break>```<break><break>`

        :param str resource_name: The name of the resource.
        :param ReplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ReplicationArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 deletion: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dest_namespace: Optional[pulumi.Input[str]] = None,
                 dest_namespace_replace: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicationFilterArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override: Optional[pulumi.Input[bool]] = None,
                 registry_id: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 speed: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReplicationArgs.__new__(ReplicationArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["deletion"] = deletion
            __props__.__dict__["description"] = description
            __props__.__dict__["dest_namespace"] = dest_namespace
            __props__.__dict__["dest_namespace_replace"] = dest_namespace_replace
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["filters"] = filters
            __props__.__dict__["name"] = name
            __props__.__dict__["override"] = override
            if registry_id is None and not opts.urn:
                raise TypeError("Missing required property 'registry_id'")
            __props__.__dict__["registry_id"] = registry_id
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["speed"] = speed
            __props__.__dict__["replication_policy_id"] = None
        super(Replication, __self__).__init__(
            'harbor:index/replication:Replication',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            deletion: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dest_namespace: Optional[pulumi.Input[str]] = None,
            dest_namespace_replace: Optional[pulumi.Input[int]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ReplicationFilterArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            override: Optional[pulumi.Input[bool]] = None,
            registry_id: Optional[pulumi.Input[int]] = None,
            replication_policy_id: Optional[pulumi.Input[int]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            speed: Optional[pulumi.Input[int]] = None) -> 'Replication':
        """
        Get an existing Replication resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReplicationState.__new__(_ReplicationState)

        __props__.__dict__["action"] = action
        __props__.__dict__["deletion"] = deletion
        __props__.__dict__["description"] = description
        __props__.__dict__["dest_namespace"] = dest_namespace
        __props__.__dict__["dest_namespace_replace"] = dest_namespace_replace
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["filters"] = filters
        __props__.__dict__["name"] = name
        __props__.__dict__["override"] = override
        __props__.__dict__["registry_id"] = registry_id
        __props__.__dict__["replication_policy_id"] = replication_policy_id
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["speed"] = speed
        return Replication(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def deletion(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "deletion")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destNamespace")
    def dest_namespace(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dest_namespace")

    @property
    @pulumi.getter(name="destNamespaceReplace")
    def dest_namespace_replace(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "dest_namespace_replace")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional[Sequence['outputs.ReplicationFilter']]]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def override(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "override")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter(name="replicationPolicyId")
    def replication_policy_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "replication_policy_id")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def speed(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "speed")

