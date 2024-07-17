import { describe, it, expect } from 'vitest';

import { mount } from '@vue/test-utils';
import FormGroup from '../../Form/FormGroup.vue';

describe('FormGroup', () => {
  it('renders the label properly', () => {
    const wrapper = mount(FormGroup, { props: { label: 'Hello Vitest' } });
    expect(wrapper.text()).toContain('Hello Vitest');
  });

  it('renders an error if supplied', () => {
    const wrapper = mount(FormGroup, { props: { label: 'Hello Vitest', error: 'I am error' } });
    expect(wrapper.text()).toContain('I am error');
  });

  it('matches snapshot', () => {
    const wrapper = mount(FormGroup, { props: { label: 'Hello Vitest' } });
    expect(wrapper.html()).toMatchSnapshot();
  });

  it('renders a label element', () => {
    const wrapper = mount(FormGroup, { props: { label: 'Hello Vitest' } });
    expect(wrapper.find('label').exists()).toBe(true);
  });

  it('renders an error message element if error is supplied', () => {
    const wrapper = mount(FormGroup, { props: { label: 'Hello Vitest', error: 'I am error' } });
    expect(wrapper.find('.error').exists()).toBe(true);
  });

  it('handles missing error prop gracefully', () => {
    const wrapper = mount(FormGroup, { props: { label: 'Hello Vitest' } });
    expect(wrapper.find('.error-message').exists()).toBe(false);
  });

  it('handles empty label prop', () => {
    const wrapper = mount(FormGroup, { props: { label: '' } });
    expect(wrapper.find('label').text()).toBe('');
  });
});
