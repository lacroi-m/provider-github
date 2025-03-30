package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func (in *WebhookParameters) DeepCopyInto(out *WebhookParameters) {
	*out = *in
	in.ConfigurableField = out.ConfigurableField
}

func (in *WebhookObservation) DeepCopyInto(out *WebhookObservation) {
	*out = *in
	in.ObservableField = out.ObservableField
}

func (in *WebhookSpec) DeepCopyInto(out *WebhookSpec) {
	*out = *in
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

func (in *WebhookStatus) DeepCopyInto(out *WebhookStatus) {
	*out = *in
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

func (in *Webhook) DeepCopyInto(out *Webhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}
func (in *Webhook) DeepCopy() *Webhook {
	if in == nil {
		return nil
	}
	out := new(Webhook)
	in.DeepCopyInto(out)
	return out
}

func (in *Webhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *WebhookList) DeepCopyInto(out *WebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Webhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *WebhookList) DeepCopy() *WebhookList {
	if in == nil {
		return nil
	}
	out := new(WebhookList)
	in.DeepCopyInto(out)
	return out
}

func (in *WebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
